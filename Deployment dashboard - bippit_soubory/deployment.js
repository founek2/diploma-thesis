import { isVersionGreater, buildTag, toDate, insertAfter } from "./util.js"
import { data } from "./data.js"

function onClickDiff(diff) {
    return (e) => {
        e.preventDefault();
        const { target } = e;
        const row = target.parentElement.parentElement;

        if (row.classList.contains("active")) {
            document.querySelectorAll(`.active ~ .diff`).forEach(e => e.remove());
        } else if (diff?.length > 0) {
            const rows = diff.map(d => buildTag("tr", [
                buildTag("td", [
                    buildTag("a", d.message, { href: d.link, target: "_blank" }),
                ], { class: "mdl-data-table__cell--non-numeric", style: "padding-left: 3rem;" }),
                buildTag("td", ""),
                buildTag("td", toDate(d.date))
            ], { class: "diff" }))

            insertAfter(rows, row)
        } else {
            const el = buildTag("tr", [
                buildTag("td", "Unknown diff, check manually"),
                buildTag("td", ""),
                buildTag("td", ""),
            ], { class: "diff" })
            insertAfter(el, row)
        }

        row.classList.toggle("active")
    }
}

const hasVersionAndTags = data => data.prodVersion && data.tags

const rows = Object.entries(data.services)
    .filter(([_, data]) => hasVersionAndTags(data))
    .map(([serviceName, { prodVersion, tags }]) =>
        tags
            .filter((tag) => isVersionGreater(tag.name, prodVersion))
            .map(tag => buildTag("tr", [
                buildTag("td", [
                    buildTag("a", serviceName, { href: "#", actions: { click: onClickDiff(tag.diff) } })
                ], { class: "mdl-data-table__cell--non-numeric" }),
                buildTag("td", [
                    buildTag("a", tag.name.replace("v", ""), { href: tag.link, target: "_blank" })
                ]),
                buildTag("td", toDate(tag.date)),
            ]))
    )

const tableElements = [
    buildTag("thead", [
        buildTag("tr", [
            buildTag("th", "Service", { class: "mdl-data-table__cell--non-numeric" }),
            buildTag("th", "Version"),
            buildTag("th", "Date"),
        ]),
    ]),
    buildTag("tbody", rows)
]

document.querySelector("table").append(...tableElements)
