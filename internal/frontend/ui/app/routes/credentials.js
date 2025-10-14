import Route from '@ember/routing/route';

export default class CredentialsRoute extends Route {
  async setupController(controller) {
    super.setupController(...arguments);
    await controller.fetchCredentials();
  }
}
