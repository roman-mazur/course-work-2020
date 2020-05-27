# course-work-2020

**<h1>Тема: Реализация ряда абстракций на языке Golang. Event Emitter, Promise,Thenable</h1>**
**<h1>Идея</h1>**

<h3>Принести в язык golang новые абстракции(в ознакомительных целях)</h3>

**<h1>Event Emitter</h1>**

<h3>За основу брал EventEmitter с nodejs пакета events. Поглядывал на <a href="https://nodejs.org/api/events.html">api</a></h3>

**<h2>Methods list</h2>**

- [GetMaxListeners](event-emitter/EVENT_EMITTER.md#h2method-getmaxlistenersh2)
- [SetMaxListeners](event-emitter/EVENT_EMITTER.md#h2method-setmaxlistenersh2)
- [Once](event-emitter/EVENT_EMITTER.md#h2method-onceh2)
- [On](event-emitter/EVENT_EMITTER.md#h2method-onh2)
- [EventNames](event-emitter/EVENT_EMITTER.md#h2method-eventnamesh2)
- [ListenerCount](event-emitter/EVENT_EMITTER.md#h2method-listenercounth2)
- [RemoveAllListeners](event-emitter/EVENT_EMITTER.md#h2method-removealllistenersh2)
- [RemoveListener](event-emitter/EVENT_EMITTER.md#h2method-removelistenerh2)
- [Emit](event-emitter/EVENT_EMITTER.md#h2method-emith2)
- [NewEventEmitter](event-emitter/EVENT_EMITTER.md#h2method-neweventemitterh2)
- [AddListener](event-emitter/EVENT_EMITTER.md#h2method-addlistenerh2)

**<h1>Promise</h1>**

<h3>За основу брал Promise с js,v8. Поглядывал на <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise">mdn</a> <a href="https://chromium.googlesource.com/v8/v8/+/3.29.45/src/promise.js?autodive=0/">source</a> <a href="https://github.com/then/promise">promises</a></h3>

<h3>Важная мысль: </h3> 
<p>На реализацию промиса стоит смотреть лишь с ознакомительной точки зрения. У меня стояла задача окончательно разобраться как под капотом работает промис, и попутно сделать свою реализацию на Golang. По скольку в го используются рутины и каналы для достижения конкурентности,а асинхронности, как таковой нет, вышло громоздкая обертка над стандартными средствами и особого смысла использовать эту обертку нет :)</p>
<h2>Methods list</h2>

- [NewPromise](promise/PROMISE.md#h2method-newpromiseh2)
- [Resolve](promise/PROMISE.md#h2method-resolveh2)
- [Reject](promise/PROMISE.md#h2method-rejecth2)
- [Then](promise/PROMISE.md#h2method-thenh2)

later...
