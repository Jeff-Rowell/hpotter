import RESTAdapter from '@ember-data/adapter/rest';

export default class ApplicationAdapter extends RESTAdapter {
  namespace = 'api';

  // Override to handle the custom API response format
  pathForType(modelName) {
    return modelName + 's';
  }
}
