import { data } from "./data.js"
import { buildTag } from "./util.js";

const targets = [{
    targetName: "Replacement of mock-it with Mockall",
    library: "mockall",
    exclude: "mock-it"
}, {
    targetName: "Cargo test without docker-compose",
    library: "postgres-test-macro"
}, {
    targetName: "Simplify DB connection",
    library: "tokio-postgres-helper"
}]

function containsDependency(name, service) {
    return service.privateDependencies.some(([depName]) => name == depName)
        || service.dependencies.some(([depName]) => name == depName)
}

const rustServices = Object.values(data.services).filter(service => service.language === "rust")

const progressElements = targets.map(({ targetName, library, exclude }) => {
    const numberOfOccurences = rustServices.map(service =>
        containsDependency(library, service)
        && (!exclude || !containsDependency(exclude, service))

    ).filter(Boolean).length
    const max = Object.keys(data.services).length;

    return buildTag("div", [
        buildTag("p", [
            targetName,
            buildTag("span", `${numberOfOccurences}`)
        ]),
        buildTag("progress", [], { class: "progress", max: `${max}`, value: `${numberOfOccurences}` })
    ], { class: "task-progress" })
}
)

document.querySelector(".progress-section").append(...progressElements)