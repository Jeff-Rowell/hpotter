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
;define("hpotter-ui/components/world-map", ["exports", "@ember/component", "@glimmer/component", "@ember/object", "@glimmer/tracking", "globe.gl", "@ember/template-factory"], function (_exports, _component, _component2, _object, _tracking, _globe, _templateFactory) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  var _class, _descriptor;
  0; //eaimeta@70e063a35619d71f0,"@glimmer/component",0,"@ember/object",0,"@glimmer/tracking",0,"globe.gl",0,"@ember/template-factory",0,"@ember/component"eaimeta@70e063a35619d71f
  function _initializerDefineProperty(e, i, r, l) { r && Object.defineProperty(e, i, { enumerable: r.enumerable, configurable: r.configurable, writable: r.writable, value: r.initializer ? r.initializer.call(l) : void 0 }); }
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  function _applyDecoratedDescriptor(i, e, r, n, l) { var a = {}; return Object.keys(n).forEach(function (i) { a[i] = n[i]; }), a.enumerable = !!a.enumerable, a.configurable = !!a.configurable, ("value" in a || a.initializer) && (a.writable = !0), a = r.slice().reverse().reduce(function (r, n) { return n(i, e, r) || r; }, a), l && void 0 !== a.initializer && (a.value = a.initializer ? a.initializer.call(l) : void 0, a.initializer = void 0), void 0 === a.initializer ? (Object.defineProperty(i, e, a), null) : a; }
  function _initializerWarningHelper(r, e) { throw Error("Decorating class property failed. Please ensure that transform-class-properties is enabled and runs after the decorators transform."); }
  const __COLOCATED_TEMPLATE__ = (0, _templateFactory.createTemplateFactory)(
  /*
    <div class="globe-wrapper">
    <div
      class="world-globe"
      {{did-insert this.setupGlobe}}
      {{will-destroy this.teardownGlobe}}
    ></div>
  
    {{#if this.selectedConnection}}
      <div class="connection-detail-panel">
        <div class="detail-header">
          <h3>Connection Details</h3>
          <button class="close-btn" {{on "click" this.closeDetails}}>×</button>
        </div>
        <div class="detail-content">
          <div class="detail-row">
            <span class="detail-label">ID:</span>
            <span class="detail-value">{{this.selectedConnection.id}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Source Address:</span>
            <span class="detail-value">{{this.selectedConnection.source_address}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Source Port:</span>
            <span class="detail-value">{{this.selectedConnection.source_port}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Destination Address:</span>
            <span class="detail-value">{{this.selectedConnection.destination_address}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Destination Port:</span>
            <span class="detail-value">{{this.selectedConnection.destination_port}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Container:</span>
            <span class="detail-value">{{this.selectedConnection.container}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Location:</span>
            <span class="detail-value">{{this.selectedConnection.lat}}, {{this.selectedConnection.lng}}</span>
          </div>
          <div class="detail-row">
            <span class="detail-label">Timestamp:</span>
            <span class="detail-value">{{this.selectedConnection.created_at}}</span>
          </div>
        </div>
      </div>
    {{/if}}
  </div>
  
  */
  {
    "id": "UBaSvpap",
    "block": "[[[10,0],[14,0,\"globe-wrapper\"],[12],[1,\"\\n  \"],[11,0],[24,0,\"world-globe\"],[4,[38,1],[[30,0,[\"setupGlobe\"]]],null],[4,[38,2],[[30,0,[\"teardownGlobe\"]]],null],[12],[13],[1,\"\\n\\n\"],[41,[30,0,[\"selectedConnection\"]],[[[1,\"    \"],[10,0],[14,0,\"connection-detail-panel\"],[12],[1,\"\\n      \"],[10,0],[14,0,\"detail-header\"],[12],[1,\"\\n        \"],[10,\"h3\"],[12],[1,\"Connection Details\"],[13],[1,\"\\n        \"],[11,\"button\"],[24,0,\"close-btn\"],[4,[38,6],[\"click\",[30,0,[\"closeDetails\"]]],null],[12],[1,\"×\"],[13],[1,\"\\n      \"],[13],[1,\"\\n      \"],[10,0],[14,0,\"detail-content\"],[12],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"ID:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"id\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Source Address:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"source_address\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Source Port:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"source_port\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Destination Address:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"destination_address\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Destination Port:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"destination_port\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Container:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"container\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Location:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"lat\"]]],[1,\", \"],[1,[30,0,[\"selectedConnection\",\"lng\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-row\"],[12],[1,\"\\n          \"],[10,1],[14,0,\"detail-label\"],[12],[1,\"Timestamp:\"],[13],[1,\"\\n          \"],[10,1],[14,0,\"detail-value\"],[12],[1,[30,0,[\"selectedConnection\",\"created_at\"]]],[13],[1,\"\\n        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n    \"],[13],[1,\"\\n\"]],[]],null],[13],[1,\"\\n\"]],[],false,[\"div\",\"did-insert\",\"will-destroy\",\"if\",\"h3\",\"button\",\"on\",\"span\"]]",
    "moduleName": "hpotter-ui/components/world-map.hbs",
    "isStrictMode": false
  });
  let WorldMapComponent = _exports.default = (_class = class WorldMapComponent extends _component2.default {
    constructor(...args) {
      super(...args);
      _defineProperty(this, "globe", null);
      _initializerDefineProperty(this, "selectedConnection", _descriptor, this);
    }
    setupGlobe(element) {
      const {
        connections
      } = this.args;

      // Create globe instance
      this.globe = (0, _globe.default)()(element).globeImageUrl('//unpkg.com/three-globe/example/img/earth-blue-marble.jpg').bumpImageUrl('//unpkg.com/three-globe/example/img/earth-topology.png').backgroundImageUrl('//unpkg.com/three-globe/example/img/night-sky.png').pointOfView({
        altitude: 2.5
      }).atmosphereColor('#3a228a').atmosphereAltitude(0.25);

      // Process connections data
      if (connections && connections.length > 0) {
        const points = connections.filter(conn => conn.latitude && conn.longitude && Math.abs(conn.latitude) > 0.0001 && Math.abs(conn.longitude) > 0.0001).map(conn => ({
          lat: conn.latitude,
          lng: conn.longitude,
          size: 0.4,
          color: '#ef4444',
          source_address: conn.source_address,
          source_port: conn.source_port,
          destination_address: conn.destination_address,
          destination_port: conn.destination_port,
          container: conn.container,
          created_at: conn.created_at,
          id: conn.id
        }));

        // Add points to globe with click handler
        this.globe.pointsData(points).pointAltitude(0.01).pointRadius('size').pointColor('color').pointLabel(point => `
          <div style="background: rgba(0,0,0,0.95); padding: 10px 12px; border-radius: 6px; font-size: 13px; color: #fff; line-height: 1.6; border: 1px solid #ef4444;">
            <div style="color: #ef4444; font-weight: 600; margin-bottom: 6px; font-size: 14px;">⚠️ Connection Details</div>
            <div><strong>Source:</strong> ${point.source_address}:${point.source_port}</div>
            <div><strong>Destination:</strong> ${point.destination_address}:${point.destination_port}</div>
            <div><strong>Container:</strong> ${point.container}</div>
            <div><strong>Time:</strong> ${new Date(point.created_at).toLocaleString()}</div>
            <div style="margin-top: 8px; padding-top: 6px; border-top: 1px solid rgba(239, 68, 68, 0.3); font-size: 11px; color: #a0a4b8;">Click for more details</div>
          </div>
        `).onPointClick(point => {
          this.handlePointClick(point);
        }).onPointHover(point => {
          // Change cursor on hover
          element.style.cursor = point ? 'pointer' : 'grab';
        });

        // Disable auto-rotation - allow manual control
        this.globe.controls().autoRotate = false;
        this.globe.controls().enableZoom = true;
        this.globe.controls().enableRotate = true;
      }

      // Handle window resize
      this.handleResize = () => {
        if (this.globe) {
          this.globe.width(element.clientWidth);
          this.globe.height(element.clientHeight);
        }
      };
      window.addEventListener('resize', this.handleResize);
    }
    handlePointClick(point) {
      if (!point) return;

      // Store selected connection
      this.selectedConnection = point;

      // Zoom to the point
      this.globe.pointOfView({
        lat: point.lat,
        lng: point.lng,
        altitude: 1.5
      }, 1000); // 1 second animation
    }
    closeDetails() {
      this.selectedConnection = null;
    }
    teardownGlobe() {
      if (this.handleResize) {
        window.removeEventListener('resize', this.handleResize);
      }
      if (this.globe) {
        // Clean up Three.js resources
        this.globe._destructor();
        this.globe = null;
      }
    }
  }, _descriptor = _applyDecoratedDescriptor(_class.prototype, "selectedConnection", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return null;
    }
  }), _applyDecoratedDescriptor(_class.prototype, "setupGlobe", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "setupGlobe"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "handlePointClick", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "handlePointClick"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "closeDetails", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "closeDetails"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "teardownGlobe", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "teardownGlobe"), _class.prototype), _class);
  (0, _component.setComponentTemplate)(__COLOCATED_TEMPLATE__, WorldMapComponent);
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
;define("hpotter-ui/controllers/connections", ["exports", "@ember/controller", "@glimmer/tracking", "@ember/object", "fetch"], function (_exports, _controller, _tracking, _object, _fetch) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  var _class, _descriptor, _descriptor2, _descriptor3, _descriptor4, _descriptor5;
  0; //eaimeta@70e063a35619d71f0,"@ember/controller",0,"@glimmer/tracking",0,"@ember/object",0,"fetch"eaimeta@70e063a35619d71f
  function _initializerDefineProperty(e, i, r, l) { r && Object.defineProperty(e, i, { enumerable: r.enumerable, configurable: r.configurable, writable: r.writable, value: r.initializer ? r.initializer.call(l) : void 0 }); }
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  function _applyDecoratedDescriptor(i, e, r, n, l) { var a = {}; return Object.keys(n).forEach(function (i) { a[i] = n[i]; }), a.enumerable = !!a.enumerable, a.configurable = !!a.configurable, ("value" in a || a.initializer) && (a.writable = !0), a = r.slice().reverse().reduce(function (r, n) { return n(i, e, r) || r; }, a), l && void 0 !== a.initializer && (a.value = a.initializer ? a.initializer.call(l) : void 0, a.initializer = void 0), void 0 === a.initializer ? (Object.defineProperty(i, e, a), null) : a; }
  function _initializerWarningHelper(r, e) { throw Error("Decorating class property failed. Please ensure that transform-class-properties is enabled and runs after the decorators transform."); }
  let ConnectionsController = _exports.default = (_class = class ConnectionsController extends _controller.default {
    constructor(...args) {
      super(...args);
      _initializerDefineProperty(this, "connections", _descriptor, this);
      _initializerDefineProperty(this, "currentPage", _descriptor2, this);
      _initializerDefineProperty(this, "pageSize", _descriptor3, this);
      _initializerDefineProperty(this, "totalCount", _descriptor4, this);
      _initializerDefineProperty(this, "isLoading", _descriptor5, this);
      _defineProperty(this, "pageSizeOptions", [10, 25, 50]);
    }
    get totalPages() {
      return Math.ceil(this.totalCount / this.pageSize);
    }
    get offset() {
      return (this.currentPage - 1) * this.pageSize;
    }
    get hasNextPage() {
      return this.currentPage < this.totalPages;
    }
    get hasPreviousPage() {
      return this.currentPage > 1;
    }
    get disableNextPage() {
      return !this.hasNextPage;
    }
    get disablePreviousPage() {
      return !this.hasPreviousPage;
    }
    get startRecord() {
      return this.offset + 1;
    }
    get endRecord() {
      const end = this.offset + this.pageSize;
      return end > this.totalCount ? this.totalCount : end;
    }
    async fetchConnections() {
      this.isLoading = true;
      try {
        const response = await (0, _fetch.default)(`/api/connections?limit=${this.pageSize}&offset=${this.offset}`);
        if (!response.ok) {
          throw new Error('Failed to fetch connections');
        }
        const data = await response.json();
        this.connections = data || [];

        // Fetch total count
        const countResponse = await (0, _fetch.default)('/api/connections');
        if (countResponse.ok) {
          const allData = await countResponse.json();
          this.totalCount = allData.length;
        }
      } catch (error) {
        console.error('Error fetching connections:', error);
        this.connections = [];
      } finally {
        this.isLoading = false;
      }
    }
    async changePageSize(event) {
      this.pageSize = parseInt(event.target.value);
      this.currentPage = 1;
      await this.fetchConnections();
    }
    async goToPage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
        await this.fetchConnections();
      }
    }
    async nextPage() {
      if (this.hasNextPage) {
        this.currentPage++;
        await this.fetchConnections();
      }
    }
    async previousPage() {
      if (this.hasPreviousPage) {
        this.currentPage--;
        await this.fetchConnections();
      }
    }
  }, _descriptor = _applyDecoratedDescriptor(_class.prototype, "connections", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return [];
    }
  }), _descriptor2 = _applyDecoratedDescriptor(_class.prototype, "currentPage", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return 1;
    }
  }), _descriptor3 = _applyDecoratedDescriptor(_class.prototype, "pageSize", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return 10;
    }
  }), _descriptor4 = _applyDecoratedDescriptor(_class.prototype, "totalCount", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return 0;
    }
  }), _descriptor5 = _applyDecoratedDescriptor(_class.prototype, "isLoading", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return false;
    }
  }), _applyDecoratedDescriptor(_class.prototype, "changePageSize", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "changePageSize"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "goToPage", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "goToPage"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "nextPage", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "nextPage"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "previousPage", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "previousPage"), _class.prototype), _class);
});
;define("hpotter-ui/controllers/credentials", ["exports", "@ember/controller", "@glimmer/tracking", "@ember/object", "fetch"], function (_exports, _controller, _tracking, _object, _fetch) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  var _class, _descriptor, _descriptor2, _descriptor3, _descriptor4, _descriptor5;
  0; //eaimeta@70e063a35619d71f0,"@ember/controller",0,"@glimmer/tracking",0,"@ember/object",0,"fetch"eaimeta@70e063a35619d71f
  function _initializerDefineProperty(e, i, r, l) { r && Object.defineProperty(e, i, { enumerable: r.enumerable, configurable: r.configurable, writable: r.writable, value: r.initializer ? r.initializer.call(l) : void 0 }); }
  function _defineProperty(e, r, t) { return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: !0, configurable: !0, writable: !0 }) : e[r] = t, e; }
  function _toPropertyKey(t) { var i = _toPrimitive(t, "string"); return "symbol" == typeof i ? i : i + ""; }
  function _toPrimitive(t, r) { if ("object" != typeof t || !t) return t; var e = t[Symbol.toPrimitive]; if (void 0 !== e) { var i = e.call(t, r || "default"); if ("object" != typeof i) return i; throw new TypeError("@@toPrimitive must return a primitive value."); } return ("string" === r ? String : Number)(t); }
  function _applyDecoratedDescriptor(i, e, r, n, l) { var a = {}; return Object.keys(n).forEach(function (i) { a[i] = n[i]; }), a.enumerable = !!a.enumerable, a.configurable = !!a.configurable, ("value" in a || a.initializer) && (a.writable = !0), a = r.slice().reverse().reduce(function (r, n) { return n(i, e, r) || r; }, a), l && void 0 !== a.initializer && (a.value = a.initializer ? a.initializer.call(l) : void 0, a.initializer = void 0), void 0 === a.initializer ? (Object.defineProperty(i, e, a), null) : a; }
  function _initializerWarningHelper(r, e) { throw Error("Decorating class property failed. Please ensure that transform-class-properties is enabled and runs after the decorators transform."); }
  let CredentialsController = _exports.default = (_class = class CredentialsController extends _controller.default {
    constructor(...args) {
      super(...args);
      _initializerDefineProperty(this, "credentials", _descriptor, this);
      _initializerDefineProperty(this, "currentPage", _descriptor2, this);
      _initializerDefineProperty(this, "pageSize", _descriptor3, this);
      _initializerDefineProperty(this, "totalCount", _descriptor4, this);
      _initializerDefineProperty(this, "isLoading", _descriptor5, this);
      _defineProperty(this, "pageSizeOptions", [10, 25, 50]);
    }
    get totalPages() {
      return Math.ceil(this.totalCount / this.pageSize);
    }
    get offset() {
      return (this.currentPage - 1) * this.pageSize;
    }
    get hasNextPage() {
      return this.currentPage < this.totalPages;
    }
    get hasPreviousPage() {
      return this.currentPage > 1;
    }
    get disableNextPage() {
      return !this.hasNextPage;
    }
    get disablePreviousPage() {
      return !this.hasPreviousPage;
    }
    get startRecord() {
      return this.offset + 1;
    }
    get endRecord() {
      const end = this.offset + this.pageSize;
      return end > this.totalCount ? this.totalCount : end;
    }
    async fetchCredentials() {
      this.isLoading = true;
      try {
        const response = await (0, _fetch.default)(`/api/credentials?limit=${this.pageSize}&offset=${this.offset}`);
        if (!response.ok) {
          throw new Error('Failed to fetch credentials');
        }
        const data = await response.json();
        this.credentials = data || [];

        // Fetch total count
        const countResponse = await (0, _fetch.default)('/api/credentials');
        if (countResponse.ok) {
          const allData = await countResponse.json();
          this.totalCount = allData.length;
        }
      } catch (error) {
        console.error('Error fetching credentials:', error);
        this.credentials = [];
      } finally {
        this.isLoading = false;
      }
    }
    async changePageSize(event) {
      this.pageSize = parseInt(event.target.value);
      this.currentPage = 1;
      await this.fetchCredentials();
    }
    async goToPage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
        await this.fetchCredentials();
      }
    }
    async nextPage() {
      if (this.hasNextPage) {
        this.currentPage++;
        await this.fetchCredentials();
      }
    }
    async previousPage() {
      if (this.hasPreviousPage) {
        this.currentPage--;
        await this.fetchCredentials();
      }
    }
  }, _descriptor = _applyDecoratedDescriptor(_class.prototype, "credentials", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return [];
    }
  }), _descriptor2 = _applyDecoratedDescriptor(_class.prototype, "currentPage", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return 1;
    }
  }), _descriptor3 = _applyDecoratedDescriptor(_class.prototype, "pageSize", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return 10;
    }
  }), _descriptor4 = _applyDecoratedDescriptor(_class.prototype, "totalCount", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return 0;
    }
  }), _descriptor5 = _applyDecoratedDescriptor(_class.prototype, "isLoading", [_tracking.tracked], {
    configurable: true,
    enumerable: true,
    writable: true,
    initializer: function () {
      return false;
    }
  }), _applyDecoratedDescriptor(_class.prototype, "changePageSize", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "changePageSize"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "goToPage", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "goToPage"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "nextPage", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "nextPage"), _class.prototype), _applyDecoratedDescriptor(_class.prototype, "previousPage", [_object.action], Object.getOwnPropertyDescriptor(_class.prototype, "previousPage"), _class.prototype), _class);
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
;define("hpotter-ui/modifiers/did-insert", ["exports", "@ember/render-modifiers/modifiers/did-insert"], function (_exports, _didInsert) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _didInsert.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/render-modifiers/modifiers/did-insert"eaimeta@70e063a35619d71f
});
;define("hpotter-ui/modifiers/did-update", ["exports", "@ember/render-modifiers/modifiers/did-update"], function (_exports, _didUpdate) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _didUpdate.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/render-modifiers/modifiers/did-update"eaimeta@70e063a35619d71f
});
;define("hpotter-ui/modifiers/will-destroy", ["exports", "@ember/render-modifiers/modifiers/will-destroy"], function (_exports, _willDestroy) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  Object.defineProperty(_exports, "default", {
    enumerable: true,
    get: function () {
      return _willDestroy.default;
    }
  });
  0; //eaimeta@70e063a35619d71f0,"@ember/render-modifiers/modifiers/will-destroy"eaimeta@70e063a35619d71f
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
    this.route('connection', {
      path: '/connection/:connection_id'
    });
    this.route('credentials');
  });
});
;define("hpotter-ui/routes/connection", ["exports", "@ember/routing/route", "fetch"], function (_exports, _route, _fetch) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/routing/route",0,"fetch"eaimeta@70e063a35619d71f
  class ConnectionRoute extends _route.default {
    async model(params) {
      try {
        const response = await (0, _fetch.default)(`/api/connection?id=${params.connection_id}`);
        if (!response.ok) {
          throw new Error('Failed to fetch connection details');
        }
        const data = await response.json();
        return data;
      } catch (error) {
        console.error('Error fetching connection details:', error);
        return null;
      }
    }
  }
  _exports.default = ConnectionRoute;
});
;define("hpotter-ui/routes/connections", ["exports", "@ember/routing/route"], function (_exports, _route) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/routing/route"eaimeta@70e063a35619d71f
  class ConnectionsRoute extends _route.default {
    async setupController(controller) {
      super.setupController(...arguments);
      await controller.fetchConnections();
    }
  }
  _exports.default = ConnectionsRoute;
});
;define("hpotter-ui/routes/credentials", ["exports", "@ember/routing/route"], function (_exports, _route) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/routing/route"eaimeta@70e063a35619d71f
  class CredentialsRoute extends _route.default {
    async setupController(controller) {
      super.setupController(...arguments);
      await controller.fetchCredentials();
    }
  }
  _exports.default = CredentialsRoute;
});
;define("hpotter-ui/routes/index", ["exports", "@ember/routing/route", "fetch"], function (_exports, _route, _fetch) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/routing/route",0,"fetch"eaimeta@70e063a35619d71f
  class IndexRoute extends _route.default {
    async model() {
      // Fetch geo data from the API
      const response = await (0, _fetch.default)('/api/geo-data?limit=1000');
      if (!response.ok) {
        throw new Error('Failed to fetch geo data');
      }
      return response.json();
    }
  }
  _exports.default = IndexRoute;
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
        <LinkTo @route="credentials">Credentials</LinkTo>
      </nav>
    </header>
  
    <main class="app-content">
      {{outlet}}
    </main>
  </div>
  
  */
  {
    "id": "O89zv4Bf",
    "block": "[[[10,0],[14,0,\"app-container\"],[12],[1,\"\\n  \"],[10,\"header\"],[14,0,\"app-header\"],[12],[1,\"\\n    \"],[10,\"h1\"],[12],[1,\"HPotter - Honeypot Monitor\"],[13],[1,\"\\n    \"],[10,\"nav\"],[12],[1,\"\\n      \"],[8,[39,4],null,[[\"@route\"],[\"index\"]],[[\"default\"],[[[[1,\"Home\"]],[]]]]],[1,\"\\n      \"],[8,[39,4],null,[[\"@route\"],[\"connections\"]],[[\"default\"],[[[[1,\"Connections\"]],[]]]]],[1,\"\\n      \"],[8,[39,4],null,[[\"@route\"],[\"credentials\"]],[[\"default\"],[[[[1,\"Credentials\"]],[]]]]],[1,\"\\n    \"],[13],[1,\"\\n  \"],[13],[1,\"\\n\\n  \"],[10,\"main\"],[14,0,\"app-content\"],[12],[1,\"\\n    \"],[46,[28,[37,7],null,null],null,null,null],[1,\"\\n  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[],false,[\"div\",\"header\",\"h1\",\"nav\",\"link-to\",\"main\",\"component\",\"-outlet\"]]",
    "moduleName": "hpotter-ui/templates/application.hbs",
    "isStrictMode": false
  });
});
;define("hpotter-ui/templates/connection", ["exports", "@ember/template-factory"], function (_exports, _templateFactory) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/template-factory"eaimeta@70e063a35619d71f
  var _default = _exports.default = (0, _templateFactory.createTemplateFactory)(
  /*
    <div class="connection-detail-page">
    {{#if @model}}
      <div class="detail-header-section">
        <h2>Connection Details - ID {{@model.id}}</h2>
        <LinkTo @route="connections" class="back-link">← Back to Connections</LinkTo>
      </div>
  
      <div class="detail-sections">
        <div class="detail-card">
          <h3>Connection Information</h3>
          <div class="detail-grid">
            <div class="detail-item">
              <span class="label">ID:</span>
              <span class="value">{{@model.id}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Created At:</span>
              <span class="value">{{@model.created_at}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Source Address:</span>
              <span class="value">{{@model.source_address}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Source Port:</span>
              <span class="value">{{@model.source_port}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Destination Address:</span>
              <span class="value">{{@model.destination_address}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Destination Port:</span>
              <span class="value">{{@model.destination_port}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Container:</span>
              <span class="value">{{@model.container}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Protocol:</span>
              <span class="value">{{@model.proto}}</span>
            </div>
            <div class="detail-item">
              <span class="label">Location:</span>
              <span class="value">
                {{#if @model.latitude}}
                  {{@model.latitude}}, {{@model.longitude}}
                {{else}}
                  Not available
                {{/if}}
              </span>
            </div>
          </div>
        </div>
  
        <div class="detail-card credentials-card">
          <h3>Captured Credentials</h3>
          {{#if @model.credentials}}
            {{#if @model.credentials.length}}
              <table class="credentials-detail-table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Username</th>
                    <th>Password</th>
                  </tr>
                </thead>
                <tbody>
                  {{#each @model.credentials as |credential|}}
                    <tr>
                      <td>{{credential.id}}</td>
                      <td><span class="username">{{credential.username}}</span></td>
                      <td><span class="password">{{credential.password}}</span></td>
                    </tr>
                  {{/each}}
                </tbody>
              </table>
            {{else}}
              <p class="no-credentials">No credentials captured for this connection.</p>
            {{/if}}
          {{else}}
            <p class="no-credentials">No credentials captured for this connection.</p>
          {{/if}}
        </div>
  
        <div class="detail-card data-card">
          <h3>Payload Data</h3>
          {{#if @model.data}}
            {{#if @model.data.length}}
              <div class="data-entries">
                {{#each @model.data as |dataEntry|}}
                  <div class="data-entry">
                    <div class="data-header">
                      <span class="data-direction {{dataEntry.direction}}">
                        {{#if (eq dataEntry.direction "inbound")}}
                          Inbound
                        {{else}}
                          Outbound
                        {{/if}}
                      </span>
                      <span class="data-id">ID: {{dataEntry.id}}</span>
                    </div>
                    <pre class="data-content">{{dataEntry.data}}</pre>
                  </div>
                {{/each}}
              </div>
            {{else}}
              <p class="no-data">No payload data captured for this connection.</p>
            {{/if}}
          {{else}}
            <p class="no-data">No payload data captured for this connection.</p>
          {{/if}}
        </div>
      </div>
    {{else}}
      <div class="error-message">
        <h2>Connection Not Found</h2>
        <p>The connection you're looking for doesn't exist or couldn't be loaded.</p>
        <LinkTo @route="connections" class="back-link">← Back to Connections</LinkTo>
      </div>
    {{/if}}
  </div>
  
  */
  {
    "id": "LJisNgZD",
    "block": "[[[10,0],[14,0,\"connection-detail-page\"],[12],[1,\"\\n\"],[41,[30,1],[[[1,\"    \"],[10,0],[14,0,\"detail-header-section\"],[12],[1,\"\\n      \"],[10,\"h2\"],[12],[1,\"Connection Details - ID \"],[1,[30,1,[\"id\"]]],[13],[1,\"\\n      \"],[8,[39,3],[[24,0,\"back-link\"]],[[\"@route\"],[\"connections\"]],[[\"default\"],[[[[1,\"← Back to Connections\"]],[]]]]],[1,\"\\n    \"],[13],[1,\"\\n\\n    \"],[10,0],[14,0,\"detail-sections\"],[12],[1,\"\\n      \"],[10,0],[14,0,\"detail-card\"],[12],[1,\"\\n        \"],[10,\"h3\"],[12],[1,\"Connection Information\"],[13],[1,\"\\n        \"],[10,0],[14,0,\"detail-grid\"],[12],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"ID:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"id\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Created At:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"created_at\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Source Address:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"source_address\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Source Port:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"source_port\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Destination Address:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"destination_address\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Destination Port:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"destination_port\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Container:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"container\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Protocol:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,[30,1,[\"proto\"]]],[13],[1,\"\\n          \"],[13],[1,\"\\n          \"],[10,0],[14,0,\"detail-item\"],[12],[1,\"\\n            \"],[10,1],[14,0,\"label\"],[12],[1,\"Location:\"],[13],[1,\"\\n            \"],[10,1],[14,0,\"value\"],[12],[1,\"\\n\"],[41,[30,1,[\"latitude\"]],[[[1,\"                \"],[1,[30,1,[\"latitude\"]]],[1,\", \"],[1,[30,1,[\"longitude\"]]],[1,\"\\n\"]],[]],[[[1,\"                Not available\\n\"]],[]]],[1,\"            \"],[13],[1,\"\\n          \"],[13],[1,\"\\n        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n\\n      \"],[10,0],[14,0,\"detail-card credentials-card\"],[12],[1,\"\\n        \"],[10,\"h3\"],[12],[1,\"Captured Credentials\"],[13],[1,\"\\n\"],[41,[30,1,[\"credentials\"]],[[[41,[30,1,[\"credentials\",\"length\"]],[[[1,\"            \"],[10,\"table\"],[14,0,\"credentials-detail-table\"],[12],[1,\"\\n              \"],[10,\"thead\"],[12],[1,\"\\n                \"],[10,\"tr\"],[12],[1,\"\\n                  \"],[10,\"th\"],[12],[1,\"ID\"],[13],[1,\"\\n                  \"],[10,\"th\"],[12],[1,\"Username\"],[13],[1,\"\\n                  \"],[10,\"th\"],[12],[1,\"Password\"],[13],[1,\"\\n                \"],[13],[1,\"\\n              \"],[13],[1,\"\\n              \"],[10,\"tbody\"],[12],[1,\"\\n\"],[42,[28,[37,12],[[28,[37,12],[[30,1,[\"credentials\"]]],null]],null],null,[[[1,\"                  \"],[10,\"tr\"],[12],[1,\"\\n                    \"],[10,\"td\"],[12],[1,[30,2,[\"id\"]]],[13],[1,\"\\n                    \"],[10,\"td\"],[12],[10,1],[14,0,\"username\"],[12],[1,[30,2,[\"username\"]]],[13],[13],[1,\"\\n                    \"],[10,\"td\"],[12],[10,1],[14,0,\"password\"],[12],[1,[30,2,[\"password\"]]],[13],[13],[1,\"\\n                  \"],[13],[1,\"\\n\"]],[2]],null],[1,\"              \"],[13],[1,\"\\n            \"],[13],[1,\"\\n\"]],[]],[[[1,\"            \"],[10,2],[14,0,\"no-credentials\"],[12],[1,\"No credentials captured for this connection.\"],[13],[1,\"\\n\"]],[]]]],[]],[[[1,\"          \"],[10,2],[14,0,\"no-credentials\"],[12],[1,\"No credentials captured for this connection.\"],[13],[1,\"\\n\"]],[]]],[1,\"      \"],[13],[1,\"\\n\\n      \"],[10,0],[14,0,\"detail-card data-card\"],[12],[1,\"\\n        \"],[10,\"h3\"],[12],[1,\"Payload Data\"],[13],[1,\"\\n\"],[41,[30,1,[\"data\"]],[[[41,[30,1,[\"data\",\"length\"]],[[[1,\"            \"],[10,0],[14,0,\"data-entries\"],[12],[1,\"\\n\"],[42,[28,[37,12],[[28,[37,12],[[30,1,[\"data\"]]],null]],null],null,[[[1,\"                \"],[10,0],[14,0,\"data-entry\"],[12],[1,\"\\n                  \"],[10,0],[14,0,\"data-header\"],[12],[1,\"\\n                    \"],[10,1],[15,0,[29,[\"data-direction \",[30,3,[\"direction\"]]]]],[12],[1,\"\\n\"],[41,[28,[37,15],[[30,3,[\"direction\"]],\"inbound\"],null],[[[1,\"                        Inbound\\n\"]],[]],[[[1,\"                        Outbound\\n\"]],[]]],[1,\"                    \"],[13],[1,\"\\n                    \"],[10,1],[14,0,\"data-id\"],[12],[1,\"ID: \"],[1,[30,3,[\"id\"]]],[13],[1,\"\\n                  \"],[13],[1,\"\\n                  \"],[10,\"pre\"],[14,0,\"data-content\"],[12],[1,[30,3,[\"data\"]]],[13],[1,\"\\n                \"],[13],[1,\"\\n\"]],[3]],null],[1,\"            \"],[13],[1,\"\\n\"]],[]],[[[1,\"            \"],[10,2],[14,0,\"no-data\"],[12],[1,\"No payload data captured for this connection.\"],[13],[1,\"\\n\"]],[]]]],[]],[[[1,\"          \"],[10,2],[14,0,\"no-data\"],[12],[1,\"No payload data captured for this connection.\"],[13],[1,\"\\n\"]],[]]],[1,\"      \"],[13],[1,\"\\n    \"],[13],[1,\"\\n\"]],[]],[[[1,\"    \"],[10,0],[14,0,\"error-message\"],[12],[1,\"\\n      \"],[10,\"h2\"],[12],[1,\"Connection Not Found\"],[13],[1,\"\\n      \"],[10,2],[12],[1,\"The connection you're looking for doesn't exist or couldn't be loaded.\"],[13],[1,\"\\n      \"],[8,[39,3],[[24,0,\"back-link\"]],[[\"@route\"],[\"connections\"]],[[\"default\"],[[[[1,\"← Back to Connections\"]],[]]]]],[1,\"\\n    \"],[13],[1,\"\\n\"]],[]]],[13],[1,\"\\n\"]],[\"@model\",\"credential\",\"dataEntry\"],false,[\"div\",\"if\",\"h2\",\"link-to\",\"h3\",\"span\",\"table\",\"thead\",\"tr\",\"th\",\"tbody\",\"each\",\"-track-array\",\"td\",\"p\",\"eq\",\"pre\"]]",
    "moduleName": "hpotter-ui/templates/connection.hbs",
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
    <div class="page-header-with-controls">
      <h2>Recent Connections</h2>
      <div class="page-size-control">
        <label for="pageSize">Show:</label>
        <select id="pageSize" {{on "change" this.changePageSize}} class="page-size-select" value={{this.pageSize}}>
          {{#each this.pageSizeOptions as |size|}}
            <option value={{size}}>{{size}}</option>
          {{/each}}
        </select>
        <span class="per-page-label">per page</span>
      </div>
    </div>
  
    <div class="connections-list">
      {{#if this.isLoading}}
        <div class="loading-indicator">
          <p>Loading connections...</p>
        </div>
      {{else if this.connections}}
        <table class="connections-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Time</th>
              <th>Source IP</th>
              <th>Source Port</th>
              <th>Destination</th>
              <th>Container</th>
              <th>Location</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {{#each this.connections as |connection|}}
              <tr>
                <td>{{connection.id}}</td>
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
                <td>
                  <LinkTo @route="connection" @model={{connection.id}} class="view-details-btn">
                    View Details
                  </LinkTo>
                </td>
              </tr>
            {{/each}}
          </tbody>
        </table>
  
        <div class="pagination-controls">
          <div class="pagination-info">
            Showing {{this.startRecord}} - {{this.endRecord}} of {{this.totalCount}} connections
          </div>
          <div class="pagination-buttons">
            <button
              type="button"
              class="pagination-btn"
              disabled={{this.disablePreviousPage}}
              {{on "click" this.previousPage}}
            >
              Previous
            </button>
            <span class="page-indicator">Page {{this.currentPage}} of {{this.totalPages}}</span>
            <button
              type="button"
              class="pagination-btn"
              disabled={{this.disableNextPage}}
              {{on "click" this.nextPage}}
            >
              Next
            </button>
          </div>
        </div>
      {{else}}
        <p>No connections found.</p>
      {{/if}}
    </div>
  </div>
  
  */
  {
    "id": "gBVbzi3m",
    "block": "[[[10,0],[14,0,\"connections-page\"],[12],[1,\"\\n  \"],[10,0],[14,0,\"page-header-with-controls\"],[12],[1,\"\\n    \"],[10,\"h2\"],[12],[1,\"Recent Connections\"],[13],[1,\"\\n    \"],[10,0],[14,0,\"page-size-control\"],[12],[1,\"\\n      \"],[10,\"label\"],[14,\"for\",\"pageSize\"],[12],[1,\"Show:\"],[13],[1,\"\\n      \"],[11,\"select\"],[24,1,\"pageSize\"],[24,0,\"page-size-select\"],[16,2,[30,0,[\"pageSize\"]]],[4,[38,4],[\"change\",[30,0,[\"changePageSize\"]]],null],[12],[1,\"\\n\"],[42,[28,[37,6],[[28,[37,6],[[30,0,[\"pageSizeOptions\"]]],null]],null],null,[[[1,\"          \"],[10,\"option\"],[15,2,[30,1]],[12],[1,[30,1]],[13],[1,\"\\n\"]],[1]],null],[1,\"      \"],[13],[1,\"\\n      \"],[10,1],[14,0,\"per-page-label\"],[12],[1,\"per page\"],[13],[1,\"\\n    \"],[13],[1,\"\\n  \"],[13],[1,\"\\n\\n  \"],[10,0],[14,0,\"connections-list\"],[12],[1,\"\\n\"],[41,[30,0,[\"isLoading\"]],[[[1,\"      \"],[10,0],[14,0,\"loading-indicator\"],[12],[1,\"\\n        \"],[10,2],[12],[1,\"Loading connections...\"],[13],[1,\"\\n      \"],[13],[1,\"\\n\"]],[]],[[[41,[30,0,[\"connections\"]],[[[1,\"      \"],[10,\"table\"],[14,0,\"connections-table\"],[12],[1,\"\\n        \"],[10,\"thead\"],[12],[1,\"\\n          \"],[10,\"tr\"],[12],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"ID\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Time\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Source IP\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Source Port\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Destination\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Container\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Location\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Actions\"],[13],[1,\"\\n          \"],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,\"tbody\"],[12],[1,\"\\n\"],[42,[28,[37,6],[[28,[37,6],[[30,0,[\"connections\"]]],null]],null],null,[[[1,\"            \"],[10,\"tr\"],[12],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"id\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"created_at\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"source_address\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"source_port\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"destination_address\"]]],[1,\":\"],[1,[30,2,[\"destination_port\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"container\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,\"\\n\"],[41,[30,2,[\"latitude\"]],[[[1,\"                  \"],[1,[30,2,[\"latitude\"]]],[1,\", \"],[1,[30,2,[\"longitude\"]]],[1,\"\\n\"]],[]],[[[1,\"                  -\\n\"]],[]]],[1,\"              \"],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,\"\\n                \"],[8,[39,17],[[24,0,\"view-details-btn\"]],[[\"@route\",\"@model\"],[\"connection\",[30,2,[\"id\"]]]],[[\"default\"],[[[[1,\"\\n                  View Details\\n                \"]],[]]]]],[1,\"\\n              \"],[13],[1,\"\\n            \"],[13],[1,\"\\n\"]],[2]],null],[1,\"        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n\\n      \"],[10,0],[14,0,\"pagination-controls\"],[12],[1,\"\\n        \"],[10,0],[14,0,\"pagination-info\"],[12],[1,\"\\n          Showing \"],[1,[30,0,[\"startRecord\"]]],[1,\" - \"],[1,[30,0,[\"endRecord\"]]],[1,\" of \"],[1,[30,0,[\"totalCount\"]]],[1,\" connections\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"pagination-buttons\"],[12],[1,\"\\n          \"],[11,\"button\"],[24,0,\"pagination-btn\"],[16,\"disabled\",[30,0,[\"disablePreviousPage\"]]],[24,4,\"button\"],[4,[38,4],[\"click\",[30,0,[\"previousPage\"]]],null],[12],[1,\"\\n            Previous\\n          \"],[13],[1,\"\\n          \"],[10,1],[14,0,\"page-indicator\"],[12],[1,\"Page \"],[1,[30,0,[\"currentPage\"]]],[1,\" of \"],[1,[30,0,[\"totalPages\"]]],[13],[1,\"\\n          \"],[11,\"button\"],[24,0,\"pagination-btn\"],[16,\"disabled\",[30,0,[\"disableNextPage\"]]],[24,4,\"button\"],[4,[38,4],[\"click\",[30,0,[\"nextPage\"]]],null],[12],[1,\"\\n            Next\\n          \"],[13],[1,\"\\n        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n\"]],[]],[[[1,\"      \"],[10,2],[12],[1,\"No connections found.\"],[13],[1,\"\\n    \"]],[]]]],[]]],[1,\"  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[\"size\",\"connection\"],false,[\"div\",\"h2\",\"label\",\"select\",\"on\",\"each\",\"-track-array\",\"option\",\"span\",\"if\",\"p\",\"table\",\"thead\",\"tr\",\"th\",\"tbody\",\"td\",\"link-to\",\"button\"]]",
    "moduleName": "hpotter-ui/templates/connections.hbs",
    "isStrictMode": false
  });
});
;define("hpotter-ui/templates/credentials", ["exports", "@ember/template-factory"], function (_exports, _templateFactory) {
  "use strict";

  Object.defineProperty(_exports, "__esModule", {
    value: true
  });
  _exports.default = void 0;
  0; //eaimeta@70e063a35619d71f0,"@ember/template-factory"eaimeta@70e063a35619d71f
  var _default = _exports.default = (0, _templateFactory.createTemplateFactory)(
  /*
    <div class="credentials-page">
    <div class="page-header-with-controls">
      <h2>Captured Credentials</h2>
      <div class="page-size-control">
        <label for="pageSize">Show:</label>
        <select id="pageSize" {{on "change" this.changePageSize}} class="page-size-select" value={{this.pageSize}}>
          {{#each this.pageSizeOptions as |size|}}
            <option value={{size}}>{{size}}</option>
          {{/each}}
        </select>
        <span class="per-page-label">per page</span>
      </div>
    </div>
  
    <div class="credentials-list">
      {{#if this.isLoading}}
        <div class="loading-indicator">
          <p>Loading credentials...</p>
        </div>
      {{else if this.credentials}}
        <table class="credentials-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Username</th>
              <th>Password</th>
              <th>Connection ID</th>
            </tr>
          </thead>
          <tbody>
            {{#each this.credentials as |credential|}}
              <tr>
                <td>{{credential.id}}</td>
                <td>
                  <span class="username">{{credential.username}}</span>
                </td>
                <td>
                  <span class="password">{{credential.password}}</span>
                </td>
                <td>
                  <LinkTo @route="connection" @model={{credential.connections_id}} class="connection-link">
                    {{credential.connections_id}}
                  </LinkTo>
                </td>
              </tr>
            {{/each}}
          </tbody>
        </table>
  
        <div class="pagination-controls">
          <div class="pagination-info">
            Showing {{this.startRecord}} - {{this.endRecord}} of {{this.totalCount}} credentials
          </div>
          <div class="pagination-buttons">
            <button
              type="button"
              class="pagination-btn"
              disabled={{this.disablePreviousPage}}
              {{on "click" this.previousPage}}
            >
              Previous
            </button>
            <span class="page-indicator">Page {{this.currentPage}} of {{this.totalPages}}</span>
            <button
              type="button"
              class="pagination-btn"
              disabled={{this.disableNextPage}}
              {{on "click" this.nextPage}}
            >
              Next
            </button>
          </div>
        </div>
      {{else}}
        <div class="no-data-message">
          <p>No credentials have been captured yet.</p>
          <p class="no-data-hint">Credentials will appear here when attackers attempt to authenticate with the honeypot.</p>
        </div>
      {{/if}}
    </div>
  </div>
  
  */
  {
    "id": "b4UP8iHC",
    "block": "[[[10,0],[14,0,\"credentials-page\"],[12],[1,\"\\n  \"],[10,0],[14,0,\"page-header-with-controls\"],[12],[1,\"\\n    \"],[10,\"h2\"],[12],[1,\"Captured Credentials\"],[13],[1,\"\\n    \"],[10,0],[14,0,\"page-size-control\"],[12],[1,\"\\n      \"],[10,\"label\"],[14,\"for\",\"pageSize\"],[12],[1,\"Show:\"],[13],[1,\"\\n      \"],[11,\"select\"],[24,1,\"pageSize\"],[24,0,\"page-size-select\"],[16,2,[30,0,[\"pageSize\"]]],[4,[38,4],[\"change\",[30,0,[\"changePageSize\"]]],null],[12],[1,\"\\n\"],[42,[28,[37,6],[[28,[37,6],[[30,0,[\"pageSizeOptions\"]]],null]],null],null,[[[1,\"          \"],[10,\"option\"],[15,2,[30,1]],[12],[1,[30,1]],[13],[1,\"\\n\"]],[1]],null],[1,\"      \"],[13],[1,\"\\n      \"],[10,1],[14,0,\"per-page-label\"],[12],[1,\"per page\"],[13],[1,\"\\n    \"],[13],[1,\"\\n  \"],[13],[1,\"\\n\\n  \"],[10,0],[14,0,\"credentials-list\"],[12],[1,\"\\n\"],[41,[30,0,[\"isLoading\"]],[[[1,\"      \"],[10,0],[14,0,\"loading-indicator\"],[12],[1,\"\\n        \"],[10,2],[12],[1,\"Loading credentials...\"],[13],[1,\"\\n      \"],[13],[1,\"\\n\"]],[]],[[[41,[30,0,[\"credentials\"]],[[[1,\"      \"],[10,\"table\"],[14,0,\"credentials-table\"],[12],[1,\"\\n        \"],[10,\"thead\"],[12],[1,\"\\n          \"],[10,\"tr\"],[12],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"ID\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Username\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Password\"],[13],[1,\"\\n            \"],[10,\"th\"],[12],[1,\"Connection ID\"],[13],[1,\"\\n          \"],[13],[1,\"\\n        \"],[13],[1,\"\\n        \"],[10,\"tbody\"],[12],[1,\"\\n\"],[42,[28,[37,6],[[28,[37,6],[[30,0,[\"credentials\"]]],null]],null],null,[[[1,\"            \"],[10,\"tr\"],[12],[1,\"\\n              \"],[10,\"td\"],[12],[1,[30,2,[\"id\"]]],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,\"\\n                \"],[10,1],[14,0,\"username\"],[12],[1,[30,2,[\"username\"]]],[13],[1,\"\\n              \"],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,\"\\n                \"],[10,1],[14,0,\"password\"],[12],[1,[30,2,[\"password\"]]],[13],[1,\"\\n              \"],[13],[1,\"\\n              \"],[10,\"td\"],[12],[1,\"\\n                \"],[8,[39,17],[[24,0,\"connection-link\"]],[[\"@route\",\"@model\"],[\"connection\",[30,2,[\"connections_id\"]]]],[[\"default\"],[[[[1,\"\\n                  \"],[1,[30,2,[\"connections_id\"]]],[1,\"\\n                \"]],[]]]]],[1,\"\\n              \"],[13],[1,\"\\n            \"],[13],[1,\"\\n\"]],[2]],null],[1,\"        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n\\n      \"],[10,0],[14,0,\"pagination-controls\"],[12],[1,\"\\n        \"],[10,0],[14,0,\"pagination-info\"],[12],[1,\"\\n          Showing \"],[1,[30,0,[\"startRecord\"]]],[1,\" - \"],[1,[30,0,[\"endRecord\"]]],[1,\" of \"],[1,[30,0,[\"totalCount\"]]],[1,\" credentials\\n        \"],[13],[1,\"\\n        \"],[10,0],[14,0,\"pagination-buttons\"],[12],[1,\"\\n          \"],[11,\"button\"],[24,0,\"pagination-btn\"],[16,\"disabled\",[30,0,[\"disablePreviousPage\"]]],[24,4,\"button\"],[4,[38,4],[\"click\",[30,0,[\"previousPage\"]]],null],[12],[1,\"\\n            Previous\\n          \"],[13],[1,\"\\n          \"],[10,1],[14,0,\"page-indicator\"],[12],[1,\"Page \"],[1,[30,0,[\"currentPage\"]]],[1,\" of \"],[1,[30,0,[\"totalPages\"]]],[13],[1,\"\\n          \"],[11,\"button\"],[24,0,\"pagination-btn\"],[16,\"disabled\",[30,0,[\"disableNextPage\"]]],[24,4,\"button\"],[4,[38,4],[\"click\",[30,0,[\"nextPage\"]]],null],[12],[1,\"\\n            Next\\n          \"],[13],[1,\"\\n        \"],[13],[1,\"\\n      \"],[13],[1,\"\\n\"]],[]],[[[1,\"      \"],[10,0],[14,0,\"no-data-message\"],[12],[1,\"\\n        \"],[10,2],[12],[1,\"No credentials have been captured yet.\"],[13],[1,\"\\n        \"],[10,2],[14,0,\"no-data-hint\"],[12],[1,\"Credentials will appear here when attackers attempt to authenticate with the honeypot.\"],[13],[1,\"\\n      \"],[13],[1,\"\\n    \"]],[]]]],[]]],[1,\"  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[\"size\",\"credential\"],false,[\"div\",\"h2\",\"label\",\"select\",\"on\",\"each\",\"-track-array\",\"option\",\"span\",\"if\",\"p\",\"table\",\"thead\",\"tr\",\"th\",\"tbody\",\"td\",\"link-to\",\"button\"]]",
    "moduleName": "hpotter-ui/templates/credentials.hbs",
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
    <div class="map-page">
    <div class="map-header">
      <h2>Connection Map</h2>
      <p class="map-info">
        {{#if @model}}
          Showing {{@model.length}} connection(s) with geographic data
        {{else}}
          No geographic data available
        {{/if}}
      </p>
    </div>
  
    <div class="map-container">
      {{#if @model}}
        <WorldMap @connections={{@model}} />
      {{else}}
        <div class="no-data">
          <p>No connections with location data found.</p>
        </div>
      {{/if}}
    </div>
  </div>
  
  */
  {
    "id": "tLRkw1H/",
    "block": "[[[10,0],[14,0,\"map-page\"],[12],[1,\"\\n  \"],[10,0],[14,0,\"map-header\"],[12],[1,\"\\n    \"],[10,\"h2\"],[12],[1,\"Connection Map\"],[13],[1,\"\\n    \"],[10,2],[14,0,\"map-info\"],[12],[1,\"\\n\"],[41,[30,1],[[[1,\"        Showing \"],[1,[30,1,[\"length\"]]],[1,\" connection(s) with geographic data\\n\"]],[]],[[[1,\"        No geographic data available\\n\"]],[]]],[1,\"    \"],[13],[1,\"\\n  \"],[13],[1,\"\\n\\n  \"],[10,0],[14,0,\"map-container\"],[12],[1,\"\\n\"],[41,[30,1],[[[1,\"      \"],[8,[39,4],null,[[\"@connections\"],[[30,1]]],null],[1,\"\\n\"]],[]],[[[1,\"      \"],[10,0],[14,0,\"no-data\"],[12],[1,\"\\n        \"],[10,2],[12],[1,\"No connections with location data found.\"],[13],[1,\"\\n      \"],[13],[1,\"\\n\"]],[]]],[1,\"  \"],[13],[1,\"\\n\"],[13],[1,\"\\n\"]],[\"@model\"],false,[\"div\",\"h2\",\"p\",\"if\",\"world-map\"]]",
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
        
