# Motivace

Dlouho jsem přemýšlel nad tématem MicroServis, protože se o jich všude mluví a každý je chce. Ovšem z praxe dostávám spíše pocit, že se tato architektura využívá až moc - zesložitěje zbytečně vývoj, tak vysoká granularita škálovatelnosti s dnešním výkonem není potřeba. Mluvím hlavně o malých až středních projektech, kde nejsou prostředky aby každé 2 micro služby měli za sebou tým vývojářu. Pak to dopadá tak, že je 20 microslužeb, 15 repozitářů 1 vývojář který se o vše má starat. Stejně tak prostředí pro běh, které je potřeba udržovat s desítkami různých technologií - např. udržovat klustr Kubernetes není vůbec jednoduché a vyžaduje velké znalosti -> s tím se komplikuje i samostatný vývoj pro vývojáře, protože musí mít celý komplexní prostředí lokálně pro efektivní vývoj. Také Debugging je poměrně velká výzva oproti Monolitické aplikaci.

# Téma diplomové práce

V rámci diplomové práce bych se chtěl zaměřit na dnes až nadužívanou architekturu MicroService, porovnat ji s architekturou klasického Monolitu a následně se podívat na architekturu modulárního monolitu označovaného někdy jako Modulith. Věnovat se budu architektuře pro malé až střední projekty, kde podle mého názoru využití MicroService architury zbytečně prodražuje a zesložiťuje vývoj aniž by to přineslo významný přinos pro projekt.

Microservice architektura vznikla v roce 2012 a velmi rychle se rozšířila do nejrůžnějších druhé SW. Dnes je již tak populární, že při návrhu nového backendu se téměř nikdo nerozmýšlí nad vhodnou architekturou, ale vše vzniká jako MicroServices. Výhody této architekturou jsou značné a mezi největší patří: škálovatelnost a technology agnostic. Bohužel se mi zdá, že již málo kdo si uvědomuje (připouští), že tato architektura jako všechny má i své negativní vlastnosti mezi něž patří např. maintenance costs and big complexity a při návrhu pro daný projekt by se měli brát v potaz všechny vlastnosti a ne pouze automaticky slepě následovat trend MicroServices a vysoké škálovatelnosti. Toto je samozřejmně v mnoha případech opodstatěné a může být tou správnou volbou, ale to platí pro velké projekty, které mají opravdu vysoký počet uživatelů.
