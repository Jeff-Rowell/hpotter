import Component from '@glimmer/component';
import { action } from '@ember/object';
import Globe from 'globe.gl';

export default class WorldMapComponent extends Component {
  globe = null;

  @action
  setupGlobe(element) {
    const { connections } = this.args;

    // Create globe instance
    this.globe = Globe()(element)
      .globeImageUrl('//unpkg.com/three-globe/example/img/earth-blue-marble.jpg')
      .bumpImageUrl('//unpkg.com/three-globe/example/img/earth-topology.png')
      .backgroundImageUrl('//unpkg.com/three-globe/example/img/night-sky.png')
      .pointOfView({ altitude: 2.5 })
      .atmosphereColor('#3a228a')
      .atmosphereAltitude(0.25);

    // Process connections data
    if (connections && connections.length > 0) {
      const points = connections
        .filter(conn => conn.latitude && conn.longitude &&
                       Math.abs(conn.latitude) > 0.0001 &&
                       Math.abs(conn.longitude) > 0.0001)
        .map(conn => ({
          lat: conn.latitude,
          lng: conn.longitude,
          size: 0.3,
          color: '#ef4444',
          source_address: conn.source_address,
          source_port: conn.source_port,
          destination_address: conn.destination_address,
          destination_port: conn.destination_port,
          container: conn.container,
          created_at: conn.created_at
        }));

      // Add points to globe
      this.globe
        .pointsData(points)
        .pointAltitude(0.01)
        .pointRadius('size')
        .pointColor('color')
        .pointLabel(point => `
          <div style="background: rgba(0,0,0,0.9); padding: 8px; border-radius: 4px; font-size: 12px;">
            <strong style="color: #ef4444;">Connection</strong><br/>
            <strong>Source:</strong> ${point.source_address}:${point.source_port}<br/>
            <strong>Dest:</strong> ${point.destination_address}:${point.destination_port}<br/>
            <strong>Container:</strong> ${point.container}<br/>
            <strong>Time:</strong> ${new Date(point.created_at).toLocaleString()}
          </div>
        `);

      // Auto-rotate the globe
      this.globe.controls().autoRotate = true;
      this.globe.controls().autoRotateSpeed = 0.5;
    }

    // Handle window resize
    this.handleResize = () => {
      if (this.globe) {
        this.globe.width(element.clientWidth);
        this.globe.height(element.clientHeight);
      }
    };
    window.addEventListener('resize', this.handleResize);
  }

  @action
  teardownGlobe() {
    if (this.handleResize) {
      window.removeEventListener('resize', this.handleResize);
    }
    if (this.globe) {
      // Clean up Three.js resources
      this.globe._destructor();
      this.globe = null;
    }
  }
}
