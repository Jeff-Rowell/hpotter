import Route from '@ember/routing/route';
import fetch from 'fetch';

export default class CredentialsRoute extends Route {
  async model() {
    try {
      const response = await fetch('/api/credentials');
      if (!response.ok) {
        throw new Error('Failed to fetch credentials');
      }
      const data = await response.json();
      return data || [];
    } catch (error) {
      console.error('Error fetching credentials:', error);
      return [];
    }
  }
}
