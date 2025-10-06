'use strict';



;define("hpotter-ui/adapters/application", ["exports", "@ember-data/adapter/rest"], function (_exports, _rest) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember-data/adapter/rest"eaimeta@70e063a35619d71f
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  class ApplicationAdapter extends _rest.default {
    constructor(...args) {
      super(...args);
      _defineProperty(this, "namespace", 'api');
    }
    // Override to handle the custom API response format
    pathForType(modelName) {
      return modelName + 's';
    }
  }
  _exports.default = ApplicationAdapter;
});
;define("hpotter-ui/app", ["exports", "@ember/application", "ember-resolver", "ember-load-initializers", "hpotter-ui/config/environment"], function (_exports, _application, _emberResolver, _emberLoadInitializers, _environment) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/application",0,"ember-resolver",0,"ember-load-initializers",0,"hpotter-ui/config/environment"eaimeta@70e063a35619d71f
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  class App extends _application.default {
    constructor(...args) {
      super(...args);
      _defineProperty(this, "modulePrefix", _environment.default.modulePrefix);
      _defineProperty(this, "podModulePrefix", _environment.default.podModulePrefix);
      _defineProperty(this, "Resolver", _emberResolver.default);
    }
  }
  _exports.default = App;
  (0, _emberLoadInitializers.default)(App, _environment.default.modulePrefix);
});
;define("hpotter-ui/component-managers/glimmer", ["exports", "@glimmer/component/-private/ember-component-manager"], function (_exports, _emberComponentManager) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _emberComponentManager.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@glimmer/component/-private/ember-component-manager"eaimeta@70e063a35619d71f
});
;define("hpotter-ui/container-debug-adapter", ["exports", "ember-resolver/container-debug-adapter"], function (_exports, _containerDebugAdapter) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _containerDebugAdapter.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"ember-resolver/container-debug-adapter"eaimeta@70e063a35619d71f
});
;define("hpotter-ui/data-adapter", ["exports", "@ember-data/debug/data-adapter"], function (_exports, _dataAdapter) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _dataAdapter.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember-data/debug/data-adapter"eaimeta@70e063a35619d71f
});
;define("hpotter-ui/initializers/ember-data", ["exports", "@ember-data/request-utils/deprecation-support"], function (_exports, _deprecationSupport) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember-data/request-utils/deprecation-support"eaimeta@70e063a35619d71f
  /*
    This code initializes EmberData in an Ember application.
  */
  var _default = _exports.default = {
    name: 'ember-data',
    initialize(application) {
      application.registerOptionsForType('serializer', {
        singleton: false
      });
      application.registerOptionsForType('adapter', {
        singleton: false
      });
    }
  };
});
;define("hpotter-ui/models/connection", ["exports", "@ember-data/model"], function (_exports, _model) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  var _dec, _dec2, _dec3, _dec4, _dec5, _dec6, _dec7, _dec8, _dec9, _class, _descriptor, _descriptor2, _descriptor3, _descriptor4, _descriptor5, _descriptor6, _descriptor7, _descriptor8, _descriptor9;
  0; //eaimeta@70e063a35619d71f0,"@ember-data/model"eaimeta@70e063a35619d71f
  function _initializerDefineProperty(e, i, r, l) { r && Object.defineProperty(e, i, { enumerable: r.enumerable, configurable: r.configurable, writable: r.writable, value: r.initializer ? r.initializer.call(l) : void 0 }); }
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  function _applyDecoratedDescriptor(i, e, r, n, l) { var a = {}; return Object.keys(n).forEach(function (i) { a[i] = n[i]; }), a.enumerable = !!a.enumerable, a.configurable = !!a.configurable, ("value" in a || a.initializer) && (a.writable = !0), a = r.slice().reverse().reduce(function (r, n) { return n(i, e, r) || r; }, a), l && void 0 !== a.initializer && (a.value = a.initializer ? a.initializer.call(l) : void 0, a.initializer = void 0), void 0 === a.initializer ? (Object.defineProperty(i, e, a), null) : a; }
  function _initializerWarningHelper(r, e) { throw Error("Decorating class property failed. Please ensure that transform-class-properties is enabled and runs after the decorators transform."); }
  let ConnectionModel = _exports.default = (_dec = (0, _model.attr)('date'), _dec2 = (0, _model.attr)('string'), _dec3 = (0, _model.attr)('number'), _dec4 = (0, _model.attr)('string'), _dec5 = (0, _model.attr)('number'), _dec6 = (0, _model.attr)('number'), _dec7 = (0, _model.attr)('number'), _dec8 = (0, _model.attr)('string'), _dec9 = (0, _model.attr)('number'), _class = class ConnectionModel extends _model.default {
    constructor(...args) {
      super(...args);
      _initializerDefineProperty(this, "created_at", _descriptor, this);
      _initializerDefineProperty(this, "source_address", _descriptor2, this);
      _initializerDefineProperty(this, "source_port", _descriptor3, this);
      _initializerDefineProperty(this, "destination_address", _descriptor4, this);
      _initializerDefineProperty(this, "destination_port", _descriptor5, this);
      _initializerDefineProperty(this, "latitude", _descriptor6, this);
      _initializerDefineProperty(this, "longitude", _descriptor7, this);
      _initializerDefineProperty(this, "container", _descriptor8, this);
      _initializerDefineProperty(this, "proto", _descriptor9, this);
    }
  }, _descriptor = _applyDecoratedDescriptor(_class.prototype, "created_at", [_dec], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor2 = _applyDecoratedDescriptor(_class.prototype, "source_address", [_dec2], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor3 = _applyDecoratedDescriptor(_class.prototype, "source_port", [_dec3], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor4 = _applyDecoratedDescriptor(_class.prototype, "destination_address", [_dec4], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor5 = _applyDecoratedDescriptor(_class.prototype, "destination_port", [_dec5], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor6 = _applyDecoratedDescriptor(_class.prototype, "latitude", [_dec6], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor7 = _applyDecoratedDescriptor(_class.prototype, "longitude", [_dec7], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor8 = _applyDecoratedDescriptor(_class.prototype, "container", [_dec8], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _descriptor9 = _applyDecoratedDescriptor(_class.prototype, "proto", [_dec9], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _class);
});
;define("hpotter-ui/router", ["exports", "@ember/routing/router", "hpotter-ui/config/environment"], function (_exports, _router, _environment) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/routing/router",0,"hpotter-ui/config/environment"eaimeta@70e063a35619d71f
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  class Router extends _router.default {
    constructor(...args) {
      super(...args);
      _defineProperty(this, "location", _environment.default.locationType);
      _defineProperty(this, "rootURL", _environment.default.rootURL);
    }
  }
  _exports.default = Router;
  Router.map(function () {
    this.route('connections');
  });
});
;define("hpotter-ui/routes/connections", ["exports", "@ember/routing/route", "@ember/service"], function (_exports, _route, _service) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  var _class, _descriptor;
  0; //eaimeta@70e063a35619d71f0,"@ember/routing/route",0,"@ember/service"eaimeta@70e063a35619d71f
  function _initializerDefineProperty(e, i, r, l) { r && Object.defineProperty(e, i, { enumerable: r.enumerable, configurable: r.configurable, writable: r.writable, value: r.initializer ? r.initializer.call(l) : void 0 }); }
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  function _applyDecoratedDescriptor(i, e, r, n, l) { var a = {}; return Object.keys(n).forEach(function (i) { a[i] = n[i]; }), a.enumerable = !!a.enumerable, a.configurable = !!a.configurable, ("value" in a || a.initializer) && (a.writable = !0), a = r.slice().reverse().reduce(function (r, n) { return n(i, e, r) || r; }, a), l && void 0 !== a.initializer && (a.value = a.initializer ? a.initializer.call(l) : void 0, a.initializer = void 0), void 0 === a.initializer ? (Object.defineProperty(i, e, a), null) : a; }
  function _initializerWarningHelper(r, e) { throw Error("Decorating class property failed. Please ensure that transform-class-properties is enabled and runs after the decorators transform."); }
  let ConnectionsRoute = _exports.default = (_class = class ConnectionsRoute extends _route.default {
    constructor(...args) {
      super(...args);
      _initializerDefineProperty(this, "store", _descriptor, this);
    }
    async model() {
      return this.store.findAll('connection');
    }
  }, _descriptor = _applyDecoratedDescriptor(_class.prototype, "store", [_service.service], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: null
  }), _class);
});
;define("hpotter-ui/serializers/application", ["exports", "@ember-data/serializer/rest"], function (_exports, _rest) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember-data/serializer/rest"eaimeta@70e063a35619d71f
  class ApplicationSerializer extends _rest.default {
    normalizeResponse(store, primaryModelClass, payload, id, requestType) {
      // Wrap the array response in an object with the model name as key
      if (Array.isArray(payload)) {
        const modelName = primaryModelClass.modelName;
        payload = {
          [modelName + 's']: payload
        };
      }
      return super.normalizeResponse(store, primaryModelClass, payload, id, requestType);
    }
  }
  _exports.default = ApplicationSerializer;
});
;define("hpotter-ui/services/store", ["exports", "@ember/debug", "ember-data/store"], function (_exports, _debug, _store) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _store.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/debug",0,"ember-data/store"eaimeta@70e063a35619d71f
  (false && !(false) && (0, _debug.deprecate)("You are relying on ember-data auto-magically installing the store service. Use `export { default } from 'ember-data/store';` in app/services/store.js instead", false, {
    id: 'ember-data:deprecate-legacy-imports',
    for: 'ember-data',
    until: '6.0',
    since: {
      enabled: '5.2',
      available: '4.13'
    }
  }));
});
;define("hpotter-ui/templates/application", ["exports", "@ember/template-factory"], function (_exports, _templateFactory) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/template-factory"eaimeta@70e063a35619d71f
  var _default = _exports.default = (0, _templateFactory.createTemplateFactory)(
  /*
    <div class="app-container">
    <header class="app-header">
      <h1>HPotter - Honeypot Monitor</h1>
      <nav>
        <LinkTo @route="index">Home</LinkTo>
        <LinkTo @route="connections">Connections</LinkTo>
      </nav>
    </header>
  
    <main class="app-content">
      {{outlet}}
    </main>
  </div>
  
  */
  {
    "id": "bm+ebCu/",
    "block": "[[[10,0],[14,0,\"app-container\"],[12],[1,\"\\n  \"],[10,\"header\"],[14,0,\"app-header\"],[12],[1,\"\\n    \"],[10,\"h1\"],[12],[1,\"HPotter - Honeypot Monitor\"],[13],[1,\"\\n    \"],[10,\"nav\"],[12],[1,\"\\n      \"],[8,[39,4],null,[[\"@route\"],[\"index\"]],[[\"default\"],[[[[1,\"Home\"]],[]]]]],[1,\"\\n      \"],[8,[39,4],null,[[\"@route\"],[\"connections\"]],[[\"default\"],[[[[1,\"Connections\"]],[]]]]],[1,\"\\n    \"],[13],[1,\"\\n  \"],[13],[1,\"\\n\\n  \"],[10,\"main\"],[14,0,\"app-content\"],[12],[1,\"\\n    \"],[46,[28,[37,7],null,null],null,null,null],[1,\"\\n  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[],false,[\"div\",\"header\",\"h1\",\"nav\",\"link-to\",\"main\",\"component\",\"-outlet\"]]",
    "moduleName": "hpotter-ui/templates/application.hbs",
    "isStrictMode": false
  });
});
;define("hpotter-ui/templates/connections", ["exports", "@ember/template-factory"], function (_exports, _templateFactory) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/template-factory"eaimeta@70e063a35619d71f
  var _default = _exports.default = (0, _templateFactory.createTemplateFactory)(
  /*
    <div class="connections-page">
    <h2>Recent Connections</h2>
  
    <div class="connections-list">
      {{#if @model}}
        <table class="connections-table">
          <thead>
            <tr>
              <th>Time</th>
              <th>Source IP</th>
              <th>Source Port</th>
              <th>Destination</th>
              <th>Container</th>
              <th>Location</th>
            </tr>
          </thead>
          <tbody>
            {{#each @model as |connection|}}
              <tr>
                <td>{{connection.created_at}}</td>
                <td>{{connection.source_address}}</td>
                <td>{{connection.source_port}}</td>
                <td>{{connection.destination_address}}:{{connection.destination_port}}</td>
                <td>{{connection.container}}</td>
                <td>
                  {{#if connection.latitude}}
                    {{connection.latitude}}, {{connection.longitude}}
                  {{else}}
                    -
                  {{/if}}
                </td>
              </tr>
            {{/each}}
          </tbody>
        </table>
      {{else}}
        <p>No connections found.</p>
      {{/if}}
    </div>
  </div>
  
  */
  {
    "id": "K+7+NDzc",
    "block": "[[[10,0],[14,0,\"connections-page\"],[12],[1,\"\\n  \"],[10,\"h2\"],[12],[1,\"Recent Connections\"],[13],[1,\"\\n\\n  \"],[10,0],[14,0,\"connections-list\"],[12],[1,\"\\n\"],[41,[30,1],[[[1,\"      \"],[10,\"table\"],[14,0,\"connections-table\"],[12],[1,\"\\n        \"],[10,\"thead\"],[12],[1,\"\\n          \"],[10,\"tr\"],[12],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Time\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Source IP\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Source Port\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Destination\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Container\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Location\"],[13],[1,\"\\n          \"],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,\"tbody\"],[12],[1,\"\\n\"],[42,[28,[37,9],[[28,[37,9],[[30,1]],null]],null],null,[[[1,\"            \"],[10,\"tr\"],[12],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"created_at\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"source_address\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"source_port\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"destination_address\"]]],[1,\":\"],[1,[30,2,[\"destination_port\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"container\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,\"\\n\"],[41,[30,2,[\"latitude\"]],[[[1,\"                  \"],[1,[30,2,[\"latitude\"]]],[1,\", \"],[1,[30,2,[\"longitude\"]]],[1,\"\\n\"]],[]],[[[1,\"                  -\\n\"]],[]]],[1,\"              \"],[13],[1,\"\\n            \"],[13],[1,\"\\n\"]],[2]],null],[1,\"        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n\"]],[]],[[[1,\"      \"],[10,2],[12],[1,\"No connections found.\"],[13],[1,\"\\n\"]],[]]],[1,\"  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[\"@model\",\"connection\"],false,[\"div\",\"h2\",\"if\",\"table\",\"thead\",\"tr\",\"th\",\"tbody\",\"each\",\"-track-array\",\"td\",\"p\"]]",
    "moduleName": "hpotter-ui/templates/connections.hbs",
    "isStrictMode": false
  });
});
;define("hpotter-ui/templates/index", ["exports", "@ember/template-factory"], function (_exports, _templateFactory) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/template-factory"eaimeta@70e063a35619d71f
  var _default = _exports.default = (0, _templateFactory.createTemplateFactory)(
  /*
    <div class="home-page">
    <h2>Welcome to HPotter</h2>
    <p>This is a honeypot monitoring interface. View recent connections and attack attempts.</p>
  
    <div class="quick-links">
      <LinkTo @route="connections" class="btn btn-primary">View Connections</LinkTo>
    </div>
  </div>
  
  */
  {
    "id": "mSwcPYkQ",
    "block": "[[[10,0],[14,0,\"home-page\"],[12],[1,\"\\n  \"],[10,\"h2\"],[12],[1,\"Welcome to HPotter\"],[13],[1,\"\\n  \"],[10,2],[12],[1,\"This is a honeypot monitoring interface. View recent connections and attack attempts.\"],[13],[1,\"\\n\\n  \"],[10,0],[14,0,\"quick-links\"],[12],[1,\"\\n    \"],[8,[39,3],[[24,0,\"btn btn-primary\"]],[[\"@route\"],[\"connections\"]],[[\"default\"],[[[[1,\"View Connections\"]],[]]]]],[1,\"\\n  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[],false,[\"div\",\"h2\",\"p\",\"link-to\"]]",
    "moduleName": "hpotter-ui/templates/index.hbs",
    "isStrictMode": false
  });
});
;define("hpotter-ui/transforms/boolean", ["exports", "@ember/debug", "@ember-data/serializer/transform"], function (_exports, _debug, _transform) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _transform.BooleanTransform;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/debug",0,"@ember-data/serializer/transform"eaimeta@70e063a35619d71f
  (false && !(false) && (0, _debug.deprecate)("You are relying on ember-data auto-magically installing the BooleanTransform. Use `export { BooleanTransform as default } from '@ember-data/serializer/transform';` in app/transforms/boolean.js instead", false, {
    id: 'ember-data:deprecate-legacy-imports',
    for: 'ember-data',
    until: '6.0',
    since: {
      enabled: '5.2',
      available: '4.13'
    }
  }));
});
;define("hpotter-ui/transforms/date", ["exports", "@ember/debug", "@ember-data/serializer/transform"], function (_exports, _debug, _transform) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _transform.DateTransform;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/debug",0,"@ember-data/serializer/transform"eaimeta@70e063a35619d71f
  (false && !(false) && (0, _debug.deprecate)("You are relying on ember-data auto-magically installing the DateTransform. Use `export { DateTransform as default } from '@ember-data/serializer/transform';` in app/transforms/date.js instead", false, {
    id: 'ember-data:deprecate-legacy-imports',
    for: 'ember-data',
    until: '6.0',
    since: {
      enabled: '5.2',
      available: '4.13'
    }
  }));
});
;define("hpotter-ui/transforms/number", ["exports", "@ember/debug", "@ember-data/serializer/transform"], function (_exports, _debug, _transform) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _transform.NumberTransform;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/debug",0,"@ember-data/serializer/transform"eaimeta@70e063a35619d71f
  (false && !(false) && (0, _debug.deprecate)("You are relying on ember-data auto-magically installing the NumberTransform. Use `export { NumberTransform as default } from '@ember-data/serializer/transform';` in app/transforms/number.js instead", false, {
    id: 'ember-data:deprecate-legacy-imports',
    for: 'ember-data',
    until: '6.0',
    since: {
      enabled: '5.2',
      available: '4.13'
    }
  }));
});
;define("hpotter-ui/transforms/string", ["exports", "@ember/debug", "@ember-data/serializer/transform"], function (_exports, _debug, _transform) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _transform.StringTransform;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/debug",0,"@ember-data/serializer/transform"eaimeta@70e063a35619d71f
  (false && !(false) && (0, _debug.deprecate)("You are relying on ember-data auto-magically installing the StringTransform. Use `export { StringTransform as default } from '@ember-data/serializer/transform';` in app/transforms/string.js instead", false, {
    id: 'ember-data:deprecate-legacy-imports',
    for: 'ember-data',
    until: '6.0',
    since: {
      enabled: '5.2',
      available: '4.13'
    }
  }));
});
;

;define('hpotter-ui/config/environment', [], function() {
  var prefix = 'hpotter-ui';
try {
  var metaName = prefix + '/config/environment';
  var rawConfig = document.querySelector('meta[name="' + metaName + '"]').getAttribute('content');
  var config = JSON.parse(decodeURIComponent(rawConfig));

  var exports = { 'default': config };

  Object.defineProperty(exports, '__esModule', { value: true });

  return exports;
}
catch(err) {
  throw new Error('Could not read config from meta tag with name "' + metaName + '".');
}

});

;
          if (!runningTests) {
            require("hpotter-ui/app")["default"].create({});
          }
        
