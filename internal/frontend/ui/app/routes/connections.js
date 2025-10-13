import Route from '@ember/routing/route';

export default class ConnectionsRoute extends Route {
  async setupController(controller) {
    super.setupController(...arguments);
    await controller.fetchConnections();
  }
}
