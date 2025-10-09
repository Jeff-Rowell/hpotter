import Component from '@glimmer/component';
import { action } from '@ember/object';
import { tracked } from '@glimmer/tracking';
import Globe from 'globe.gl';

export default class WorldMapComponent extends Component {
  globe = null;
  @tracked selectedConnection = null;

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
          size: 0.4,
          color: '#ef4444',
          source_address: conn.source_address,
          source_port: conn.source_port,
          destination_address: conn.destination_address,
          destination_port: conn.destination_port,
          container: conn.container,
          created_at: conn.created_at,
          id: conn.id
        }));

      // Add points to globe with click handler
      this.globe
        .pointsData(points)
        .pointAltitude(0.01)
        .pointRadius('size')
        .pointColor('color')
        .pointLabel(point => `
          <div style="background: rgba(0,0,0,0.95); padding: 10px 12px; border-radius: 6px; font-size: 13px; color: #fff; line-height: 1.6; border: 1px solid #ef4444;">
            <div style="color: #ef4444; font-weight: 600; margin-bottom: 6px; font-size: 14px;">⚠️ Connection Details</div>
            <div><strong>Source:</strong> ${point.source_address}:${point.source_port}</div>
            <div><strong>Destination:</strong> ${point.destination_address}:${point.destination_port}</div>
            <div><strong>Container:</strong> ${point.container}</div>
            <div><strong>Time:</strong> ${new Date(point.created_at).toLocaleString()}</div>
            <div style="margin-top: 8px; padding-top: 6px; border-top: 1px solid rgba(239, 68, 68, 0.3); font-size: 11px; color: #a0a4b8;">Click for more details</div>
          </div>
        `)
        .onPointClick((point) => {
          this.handlePointClick(point);
        })
        .onPointHover((point) => {
          // Change cursor on hover
          element.style.cursor = point ? 'pointer' : 'grab';
        });

      // Disable auto-rotation - allow manual control
      this.globe.controls().autoRotate = false;
      this.globe.controls().enableZoom = true;
      this.globe.controls().enableRotate = true;
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
  handlePointClick(point) {
    if (!point) return;

    // Store selected connection
    this.selectedConnection = point;

    // Zoom to the point
    this.globe.pointOfView({
      lat: point.lat,
      lng: point.lng,
      altitude: 1.5
    }, 1000); // 1 second animation
  }

  @action
  closeDetails() {
    this.selectedConnection = null;
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
