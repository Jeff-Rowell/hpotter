import Model, { attr } from '@ember-data/model';

export default class ConnectionModel extends Model {
  @attr('date') created_at;
  @attr('string') source_address;
  @attr('number') source_port;
  @attr('string') destination_address;
  @attr('number') destination_port;
  @attr('number') latitude;
  @attr('number') longitude;
  @attr('string') container;
  @attr('number') proto;
}
