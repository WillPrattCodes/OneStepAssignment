<template>
  <div ref="mapContainer" class="w-full h-screen"></div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useGPSStore } from '@/stores/gps';
import { usePrefStore } from '@/stores/pref';
import { loadGoogleMaps } from '@/utils/googleMapLoader'; //import function to load google maps
import eventBus from '@/utils/eventBus'; //import event bus to listen for events

//get gps store to access gps data
const gpsStore = useGPSStore();
const prefStore = usePrefStore();
const mapContainer = ref(null);
const map = ref(null);
let google = null; //global variable to store maps object
const markers = {}; // Store markers by device_id

//load maps api key from .env file
const apiKey = import.meta.env.VITE_MAPS_API;

//get gps data using store getter (reactive)
const gpsData = computed(() => gpsStore.gpsData);
const hiddenDevices = computed(() => prefStore.preferences.hiddenDevices);

//load google maps and initialize map
const initializeMap = async () => {
  try {
    //set global google variable to maps object
    google = await loadGoogleMaps(apiKey);

    //give map default center of los angeles
    const center = { lat: 34.0522, lng: -118.2437 };
    map.value = new google.maps.Map(mapContainer.value, {
      center,
      zoom: 10,
    });

    //add gps markers for all devices
    updateMarkers(gpsData.value);
  } catch (error) {
    console.error('Failed to load Google Maps:', error);
  }
};

//add or update markers on map
const updateMarkers = (gpsData) => {
  console.log('Updating Markers:', gpsData);
  if (!google || !map.value) return; //make sure maps is loaded

  gpsData.forEach((device) => {
    const { device_id } = device;

    //check if device is hidden
    if (hiddenDevices.value.includes(device_id)) {
      //remove marker if hidden
      if (markers[device_id]) {
        markers[device_id].setMap(null);
        delete markers[device_id];
      }
      return; //skip rendering for hidden devices
    }

    if (device.latest_device_point) {
      const { lat, lng } = device.latest_device_point;

      //update existing markers
      if (markers[device_id]) {
        markers[device_id].setPosition({ lat, lng });
      } else {
        //create new markers if new devices added
        const marker = new google.maps.Marker({
          position: { lat, lng },
          map: map.value,
          title: device.display_name,
        });

        //create info window for each marker
        const infoWindow = new google.maps.InfoWindow({
          content: `<h3>${device.display_name}</h3>
                    <p>Status: ${device.online ? 'Online' : 'Offline'}</p>
                    <p>Speed: ${Math.round(device.latest_device_point.speed)}</p>`,
        });

        //when marker clicked show info window
        marker.addListener('click', () => {
          infoWindow.open(map.value, marker);
        });

        //store marker in markers object
        markers[device_id] = marker;
      }
    }
  });
};

//listen for visibility toggle event using event listener
eventBus.on('toggle-visibility', () => {
  updateMarkers(gpsData.value); //re-render markers on visibility
});

//watch for changes in gps data and hidden devices
watch(gpsData, (newData) => {
  updateMarkers(newData);
});

//initialize map on component mount
onMounted(async () => {
  await initializeMap(); //initialize google map
});
</script>
