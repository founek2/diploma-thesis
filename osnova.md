# Diplomka

Dlouho jsem přemýšlel nad tématem MicroServis, protože se o jich všude mluví a každý je chce. Ovšem z praxe dostávám spíše pocit, že se tato architektura využívá až moc - zesložitěje zbytečně vývoj, tak vysoká granularita škálovatelnosti s dnešním výkonem není potřeba. Mluvím hlavně o malých až středních projektech, kde nejsou prostředky aby každé 2 micro služby měli za sebou tým vývojářu. Pak to dopadá tak, že je 20 microslužeb, 15 repozitářů 1 vývojář který se o vše má starat. Stejně tak prostředí pro běh, které je potřeba udržovat s desítkami různých technologií - např. udržovat klustr Kubernetes není vůbec jednoduché a vyžaduje velké znalosti -> s tím se komplikuje i samostatný vývoj pro vývojáře, protože musí mít celý klastr lokálně pro efektivní vývoj.

V rámci Diplomky, bych chtěl porovnat přístupy Monolit, MicroService a modulární Monolit, který pokud se správně navrhne tak dokáže být přínosem pro menší projekty. Chtěl bych dojít k závěru, že pro malé až střední projekty by se mělo začínat s Monolitem a pokud je dobře navržený, tak je možné ho dle potřeb postupně převést na Service oriented architecture (nebo možná i microService?). Samotný vývoj a údržba Micro služeb dle mého názoru zbytečně zesložiťuje v hodně případech život programátorům a stojí firmy peníze na víc, jenom protože chtějí být cool a mít "MicroService". Samozřejmně ve spoustě případech jsou MicroService nejlepším řešení s ohledem na jejich škálovatelnost a nezávislost.

## Osnova

Jedná se čistě o nastřelený návrh.

1. Analýza monolitu
    1. snažší návrh, rychlejší vývoj
    2. problém škálovatelnosti a provázanosti
    3. ACID
2. Microservices
    1. složitý návrh ale dobrá škálovatelnost by design
    2. velké množství potřebných technologií pro spuštění jednoduché aplikace
    3. všichni dnes chtějí microservice, ale je složité je dobře navrhnout a často přinášení více problémů než užitku - minimálně pro malé/středně velké projekty
    4. Složitost distribuovaných transakcí
3. Dev methology
    1. Slightly touching Agile, since it closely relates with fast iterative changes and influences style of development
4. Modulární monolyt
    1. Monolyt s kvalitním designem a modulární architekturou může být mnohem jednodušší na vývoj a v případě nutnosti s trochou snahy umožňuje rozdělení na více služeb a díky tomu škálovat.
5. Návrh modulárního monolytu a microservice architektury
6. Porovnání potřebných technologií pro produkční deploy, cena zdrojů, využití paměti/cpu, škálovatelnost (otázka do jaké míry je tato vlastnost potřeba na drtivou většinou úloh na dnešním HW)

7. ~~Aplikace zvoleného přístupu na mojí IoT Platformu a provedení měření, zjištění chování při velkém množství zařízení -> objevení nejvytíženější částí a extrakce do samotné služby a naškálování?~~

### Nápady

-   návrh jednoduché aplikace pro všechny tři architektury na které budou dobře vidět praktické rozdíly - api, objedávka, platba, sklad
