import Route from '@ember/routing/route';
import fetch from 'fetch';

export default class MapRoute extends Route {
  async model() {
    // Fetch geo data from the API
    const response = await fetch('/api/geo-data?limit=1000');
    if (!response.ok) {
      throw new Error('Failed to fetch geo data');
    }
    return response.json();
  }
}
