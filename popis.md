# Motivace

Dlouho jsem přemýšlel nad tématem MicroServis, protože se o jich všude mluví a každý je chce. Ovšem z praxe dostávám spíše pocit, že se tato architektura využívá až moc - zesložitěje zbytečně vývoj, tak vysoká granularita škálovatelnosti s dnešním výkonem není potřeba. Mluvím hlavně o malých až středních projektech, kde nejsou prostředky aby každé 2 micro služby měli za sebou tým vývojářu. Pak to dopadá tak, že je 20 microslužeb, 15 repozitářů 1 vývojář který se o vše má starat. Stejně tak prostředí pro běh, které je potřeba udržovat s desítkami různých technologií - např. udržovat klustr Kubernetes není vůbec jednoduché a vyžaduje velké znalosti -> s tím se komplikuje i samostatný vývoj pro vývojáře, protože musí mít celý komplexní prostředí lokálně pro efektivní vývoj. Také Debugging je poměrně velká výzva oproti Monolitické aplikaci.

# Téma diplomové práce

V rámci diplomové práce bych se chtěl zaměřit na dnes až nadužívanou architekturu MicroService, porovnat ji s architekturou klasického Monolitu a následně se podívat na architekturu modulárního monolitu označovaného někdy jako Modulith. Věnovat se budu architektuře pro malé až střední projekty, kde podle mého názoru využití MicroService architury zbytečně prodražuje a zesložiťuje vývoj aniž by to přineslo významný přinos pro projekt.

Microservice architektura vznikla v roce 2012 a velmi rychle se rozšířila do nejrůžnějších druhé SW. Dnes je již tak populární, že při návrhu nového backendu se téměř nikdo nerozmýšlí nad vhodnou architekturou, ale vše vzniká jako MicroServices. Výhody této architekturou jsou značné a mezi největší patří: škálovatelnost a technology agnostic. Bohužel se mi zdá, že již málo kdo si uvědomuje (připouští), že tato architektura jako všechny má i své negativní vlastnosti mezi něž patří např. maintenance costs and big complexity a při návrhu pro daný projekt by se měli brát v potaz všechny vlastnosti a ne pouze automaticky slepě následovat trend MicroServices a vysoké škálovatelnosti. Toto je samozřejmně v mnoha případech opodstatěné a může být tou správnou volbou, ale to platí pro velké projekty, které mají opravdu vysoký počet uživatelů.

<!-- S dnešním HW již není problém mít server v konfiguraci 64 jader CPU a 0.5 TB ram, který by měl zvládnout bez větších problému libovolně složitou aplikace. -->

## Monolit

V první části práci bych se chtěl blíže podívat na klasický monolit a trochu poodkrýt představu monolitu, kterou většina lidí má: že se jedná o obrovských moloch, napsaný v Javě s množstvím závislostí, ve kterých se nikdo nevyzná a každý se bojí udělat v kódu jakoukoli změnu. Takto by samozřejmně architektura monolitu vypadat neměla. Chtěl bych tedy analyzovat architekturu Monolitu a ukázat, že některé negativní vlasti nevychájí striktně z architektury Monolitu jako takové, ale spíše ze špatného návrhu konkrétního SW - velkou nevýhodu vidím v tom, že Monolith nechává velkou volnost v návrhu a sám nevynucuje velkou míru oddělení (což se ukázalo jako nedílná součast udržitelné architektury) jako v případě např. MicroService, kde rozdělení odpovědnosti je přímo základním prvkem architektury. Tedy mnohem více rozhodnutí a volnosti zůstává v rukách SW architekta a ne vždy toto může dopadnou dobře.

Výhody:

-   Simplicity of development
-   Simplicity of debugging
-   Simplicity of deployment (one deployement unit)
-   Low cost in the early stages of the application
-   ACID

> Zde chci zdůraznit násobně rychlejší vývoj v raných fázích, mnohem nižší počet potřebných technologií a s tím související nižší náklady.

Nevýhody:

-   Slow speed of development - slower CI/CD, coordination of feature development and parallel work
-   High code coupling
-   Code ownership cannot be used
-   Performance issues
-   The cost of infrastructure
-   Lack of flexibility - tight to the technologies that are used inside your monolith

> Zde bych se chtěl více věnovat otázce výkonu a škálovatelnosti, protože to je dnes bráno jako největší negativum této architektury. V tomto ohlednu se vyvinuli různé přístupy a druhy aplikačních serverů, které se snažily tento problém řešit.

## Microservices

Dnes hojně využívaná architektura, která přinesla revoluci do způsobu jak se dnes navrhují a píší aplikace, jejímž cílem bylo vyřešit nejbolestsivější problém Monolitickým aplikací kolem škálovatelnosti a udržitelnosti velkých aplikací. O pozitivních vlastností této architektury toho bylo napsáno mnoho, ale já bych se chtěl zaměřit více na ty negativní (stinné stránky):

-   škálovatelnost - na první pohled to vypadá, že při použití této architektury bude každá aplikace libovolně škálovat a vše bude skvělé. Ovšem extrémně záleží na tom, jak jsou dané služby napsané. Do jisté míry to pravda je - při špatně napsané službě sice lze aplikaci také naškálovat a zvýšit rychlost např. rozdělit zpracování na 10 instancí, ovšem nemusí dojít ke znatelnému zrychlení protože bottleneck se přenuse jinam např. na databázi, protože aplikace vytváří příliš mnoho spojení.
-   fault isolation - v případě monolitu chyba v jedné komponentě znamená často selhání celku. U mikroslužeb by služby měli být nezávislé a chyba v jedné by neměla ovlivnit funkčnosti ostatních. To je ovšem opět závislé na konkrátní implementaci a je velmi snadné vytvořit MicroServices, které jsou na sobě zcela závislé a aplikační chyba v jedné způsobí problém v ostatních.
-   Program language and technology agnostic - volnost používat různé jazyky pro řešení problémů pro které se hodí je skvělá vlastnost. Ovšem toto přináší komplexitu, kdy vývojáři při úpravě služeb musejí znát dané jazyky a často různé jazyky používají ruzná paradigmata, takže opět vyšší nároky. Je tedy důležité mít velmi dobrý důvod toto využít.
-   Simpler to deploy - lze nasazovat nové verze služeb nezávisle na sobě. Toto je velmi dobrá vlastnost, která ale opět přináší vyšší komplexitu v psaní služeb. Je zde vyžadována maximální kompatibilita rozhraní, rozšiřovat je lze poměrně snadno ovšem problém je v odebírání funkcionalit - jak dlouho držet zpětnou kompatibilitu? Jak sledovat zda některá jiná služba dané rozhraní ještě využívá nebo jej lze bezpečně odstranit? Toto vyžaduje velkou obezřetnost a může vést k tomu, že ze strachu raději nic neodebíráme a budeme jenom rozšiřovat, což povede k udržování funkcí, které jsou naprosto zbytečné a nevyužité.
-   Complexity - Asi nejvíce používaným orchestrátorem je dnes Kubernetes. Ovšem jenom nasazení a provoz Kubernetes prostředí vyžaduje obrovské znalosti. Sice lze využít služby as a service, ale ne všem problémům se lze tímto způsobem vyhnout. Debugging v MicroServices prostředí je poměrně velká výzva. Tím, že služby běží na klusteru uzlů, tak veškerá komunikace mezi nimi nyní přidává novou komplexitu, která v Monolitu, který běží na jednom serveru nikdy nebyla - nespolehlivost komunikace, protože síť negarantuje spolehlivost, zatímco komunikace mezi procesy v rámci jednoho OS ano.
-   Distributed transactions - complexity

## Modulith

Vývoj Monolitu za dekády velmi pokročil a vyvinuli se modernější přístupy, které integrují nové poznatky z MicroService světa. Tato architektura by měla průnikem pozitivních vlastností Monolitu (rychlý vývoj, jednoduché nasazení, ACID) a z MicroServices se ukázala jako enormně důležitá vlastnost low coupling pro maintanable code. Částí kódu by měli být rozděleny do komponent, které by neměli mít mezi sebou závislost případně velmi minimální. Díky tomu by mělo být možné v budoucnu komponenty dle potřeby vyčlenit do mikroslužeb a ty škálovat.

Kód by měl být dělený do izolovaných komponent, které mezi sebou budou komunikace pomocí message service - synchroní v případě, že je potřeba dodržet ACID. Je zde otázka jak moc by měli být izolované - jestli např. i na úrovni DB, pokud ano, tak by to v podstatě byla MicroService architektura s tím, že MicroSlužby se vysazují jako jeden celek, což by mělo jisté výhody.

### Návrh nového SW

Tato architektura by mohla být novým standardem co volit pro rychlý vývoj a přitom si zachovat možnost kdykoli přejít poměrně pohodlně na MicroServices.

### Monolith to Modulith

Vytvoření postupu jak existující Monolit "učesat" a rozčlenit do komponent.

### MicroService to Modulith

Postup jak snížit počet mikro-služeb při zachování low-coupling. Výsledek by byl hybrid, kdy z mikro služeb by se stali větší služby, kde každá by měla Modulith architekturu pohlcením více mikroslužeb do jedné. Důležitým faktorem je zde správná identifikace kdy se pro tento krok rozhodnout.

# Praxe

Pracuji půl roku ve společnosti, kde máme architekturu Microservices s 20 mikroslužbami, vyvíjeno 3 roky, napsáno v Rustu, a plánuji ji použít v rámci práce jako ukázku jak se to nemá dělat. Obrovská provázanost/závislost služeb. Nic jako distribuované transakce nepoužívají - chyba znamená ruční hledání příčiny a nápravu. Spousta mikroslužeb a pouze 2 vyvojáři. Odebírání funkcí z rozhraní je čistě na vývojáří aby věděl že nic potenciálně nerozbije. Nulový monitoring u služeb jaké všechny rozhraní z ostatních služeb volá a závisí tak na nich.

## Context

Purpose of platform is to connect users with their personal financial advisor. Users can have different level of subscription and can connect their account`s to platform, which will process their transaction into nice overview.

## Example architecture

Graph below shows current Microservice architecture. Every node is MicroService and every edge represents dependency (e.g. "service1 -> service2" means service1 depends on service2). There is no API gateway and every MicroService has it`s API accessible from outside.

![Current Microservice architecture graph](_media/microservices-current-commented2.png)

Just by looking at graph you see that there is quite complexity and this contains only direct dependencies via http, it does not contain relations via async queues nor any dependencies on third-party services.

### Advantages

scaling - every service has it`s own database and service discovery is in place, so it is possible to scale every service independently.

Single Responsibility - thinking in terms of MicroServices it is forcing developer to separate unrelated features, which is very good to keep unrelated parts separated.

### Cons

scaling - In reality there is no autoscaling, just running 2 instances of each service

Http communication - http protocol is not bad by itself, but inherently it does not contain any type definition. Yes, there are some solution like Swagger, which is used in this case, but since it is not part of http itself, it is harder to manage. E.g. swagger definitions can become out of sync with code. There are some tools which can be used to automate this process, but not for every language and it is not always possible. More modern solutions like gRPC have type definition as part of it`s core and are easier to use and integrate (+ little bit faster).

no transactions - in this example transactions are not used anywhere. I presume there was presumtion not to use/care about transaction on MS level, because most of implementation contains just around 1-3 db calls, so low risk something will fail. But no transactions anywhere across whole platform is just bad design decision. There are not even revert actions, so when any error happends, it usually just results in http 500. It will be picked by alerting system and some developer should pick it up and investigate + fix it.

development speed - to add/edit feature usually means touching multiple MS. Since every http call has a lot of subsequent calls, thus lot of dependencies and it becomes quite hard to write tests, which should be as easy as possbile. This is not inherently cons of MicroService architecture, but rather of bad design.

Error prone - Since whole platform is written in Rust, which is strictly typed language, it is safe for single service. But when it comes to communication, it is imposible to enforce type safety for http calls. Every MS usually defines it`s own struct for parsing http body and there is no way how to enforce, that it was defined properly e.g. does not contain any required field, which is not returned by target API.

no failover - even though MSs architecture should be highly available by design, it always come down to the actuall implementation. In this case whener any service recieves any error from dependent service, it just propagates the error. And since there is so much dependencies, it is very likely, that whenever there will be an error in one MS (either because of programmer mistake an invalid DB data) the vast majority of platform will suddently stop working until the error is fixed - probably by reverting to older version.

> MS = MicroService

### Monolit

<!-- Vysoká provázanost -> Monolit zaručí type safety (compile time checked), potřeba škálovat pouze 3 služeb zpracovávající transakce, lze vytvoři jednoduše ACID vs. komplexita distribuovaných transakcí (koordinátor/saga), no autoscaling (HA simply by running two instances)-->

So in this case there are just two advantages of MircoService architecture:

-   scaling
-   logical separation (spliting code into multiple MS)

Converting this into modular monolith we would gain:

-   logical separation - this should be enforced by tools, ideally by compiler or run-time
-   scaling - since whole platform has couple of thousands users and is written in Rust (high performance language) every service has just 7MB memory footprint most of the time -> 20 (num. of services) x 7MB \* 3 (num. of instances) <= 500 MB, so there is no issue running this with low resources in multiple instances in terms of CPU/Memory usage
-   no need for distributed transactions - We could use just normal transactions. Or we can decide to add distributed transactions later on, but inherently they add a lot of additional complexity (e.g. Saga pattern or coordinator)
-   type safe interfaces and comminucation across whole platform at compile time

> I think this will apply to a lot of other projects with MicroService architecture as well
