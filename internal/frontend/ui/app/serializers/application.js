import RESTSerializer from '@ember-data/serializer/rest';

export default class ApplicationSerializer extends RESTSerializer {
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
