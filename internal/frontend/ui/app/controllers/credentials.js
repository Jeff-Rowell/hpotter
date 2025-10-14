import Controller from '@ember/controller';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import fetch from 'fetch';

export default class CredentialsController extends Controller {
  @tracked credentials = [];
  @tracked currentPage = 1;
  @tracked pageSize = 10;
  @tracked totalCount = 0;
  @tracked isLoading = false;

  pageSizeOptions = [10, 25, 50];

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
      const response = await fetch(`/api/credentials?limit=${this.pageSize}&offset=${this.offset}`);
      if (!response.ok) {
        throw new Error('Failed to fetch credentials');
      }
      const data = await response.json();
      this.credentials = data || [];

      // Fetch total count
      const countResponse = await fetch('/api/credentials');
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

  @action
  async changePageSize(event) {
    this.pageSize = parseInt(event.target.value);
    this.currentPage = 1;
    await this.fetchCredentials();
  }

  @action
  async goToPage(page) {
    if (page >= 1 && page <= this.totalPages) {
      this.currentPage = page;
      await this.fetchCredentials();
    }
  }

  @action
  async nextPage() {
    if (this.hasNextPage) {
      this.currentPage++;
      await this.fetchCredentials();
    }
  }

  @action
  async previousPage() {
    if (this.hasPreviousPage) {
      this.currentPage--;
      await this.fetchCredentials();
    }
  }
}
