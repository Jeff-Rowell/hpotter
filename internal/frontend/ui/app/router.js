import EmberRouter from '@ember/routing/router';
import config from 'hpotter-ui/config/environment';

export default class Router extends EmberRouter {
  location = config.locationType;
  rootURL = config.rootURL;
}

Router.map(function () {
  this.route('connections');
  this.route('connection', { path: '/connection/:connection_id' });
  this.route('map');
  this.route('credentials');
});
