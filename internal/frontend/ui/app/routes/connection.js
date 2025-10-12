import Route from '@ember/routing/route';
import fetch from 'fetch';

export default class ConnectionRoute extends Route {
  async model(params) {
    try {
      const response = await fetch(`/api/connection?id=${params.connection_id}`);
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
