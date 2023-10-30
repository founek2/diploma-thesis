import { data } from "./data.js"

function createData(includedLibraries = [], includeExternal = false) {
    const nodesData = [];
    const edgesData = [];
    const externalServices = new Set();
    for (const serviceName in data.services) {
        if (!includeExternal) break;

        const service = data.services[serviceName];
        (service.clients || []).forEach((name) => {
            if (!data.services[name]) externalServices.add(name);
        });
    }

    for (const serviceName of [...externalServices]) {
        nodesData.push({
            id: serviceName,
            label: serviceName,
            color: 'pink',
        });
    }

    for (const serviceName in data.services) {
        const service = data.services[serviceName];

        const deployedLatestVersion = service.version === service.devVersion;
        nodesData.push({
            id: serviceName,
            label: `${serviceName} ${service.version}`,
            color: !deployedLatestVersion ? 'rgb(33, 150, 243)' : 'rgb(100, 181, 246)',
            title: `dev:  ${service.devVersion}
                stag: ${service.stagVersion}
                prod: ${service.prodVersion}
                ----------
                ${service?.privateDependencies?.map(([name, version]) => `${name}: ${version}`).join('\n')}`,
        });

        for (const [depName, depVersion] of service.privateDependencies || []) {
            if (!data.libraries[depName]) continue;
            edgesData.push({ from: serviceName, to: depName, arrows: 'to', label: depVersion });
        }

        // Skip edges between services when showing libraries
        if (includedLibraries.length > 0) continue;
        for (const clientName of service.clients || []) {
            // if (!data.services[clientName]) continue;
            edgesData.push({ from: serviceName, to: clientName, arrows: 'to' });
        }
    }

    for (const libraryName in data.libraries) {
        if (!includedLibraries.includes(libraryName)) continue;

        const library = data.libraries[libraryName];
        nodesData.push({
            id: libraryName,
            label: `${libraryName} ${library.version}`,
            title: library?.privateDependencies.join('\n'),
            color: 'red',
        });
    }

    var visData = {
        nodes: new vis.DataSet(nodesData),
        edges: new vis.DataSet(edgesData),
    };

    return visData;
}

// create a network
var container = document.getElementById('mynetwork');

// provide the data in the vis format

var options = {};

// initialize your network!
const visData = createData();
var network = new vis.Network(container, visData, options);
network.on('stabilizationIterationsDone', function () {
    network.setOptions({ physics: false });
});

let savedEdges;
network.on('click', function (properties) {
    var ids = properties.nodes;
    if (savedEdges) {
        network.body.data.edges.add(savedEdges);
        savedEdges = undefined;
    } else if (ids.length > 0) {
        const ID = ids[0];

        const connectedEdgeIDs = network.getConnectedEdges(ID);
        const edgesToDisconnectIDs = visData.edges.getIds().filter((id) => !connectedEdgeIDs.includes(id));

        savedEdges = network.body.data.edges.get(edgesToDisconnectIDs);
        network.body.data.edges.remove(edgesToDisconnectIDs);
    }
});

let includedLibraries = [];
const picker = document.getElementById('library-picker');

for (const libraryName in data.libraries) {
    const input = document.createElement('input');
    input.type = 'checkbox';
    input.name = libraryName;
    input.value = libraryName;
    input.addEventListener('click', function (e) {
        if (e.target.checked) includedLibraries.push(libraryName);
        else includedLibraries = includedLibraries.filter((l) => l !== libraryName);

        network.setOptions({ physics: true });
        network.setData(createData(includedLibraries));
    });
    picker.append(input);

    const label = document.createElement('label');
    label.innerHTML = libraryName;
    picker.append(label);
}
document.getElementById('show-external-input').addEventListener('click', function (e) {
    network.setOptions({ physics: true });
    network.setData(createData(includedLibraries, e.target.checked));
});