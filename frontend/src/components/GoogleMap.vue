<template>
  <div ref="mapContainer" class="w-full h-screen"></div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useGPSStore } from '@/stores/gps';
import { loadGoogleMaps } from '@/utils/googleMapLoader'; //import function to load google maps

//get gps store to access gps data
const gpsStore = useGPSStore();
const mapContainer = ref(null);
const map = ref(null);
let google = null; //global variable to store maps object
const markers = {};

//load maps api key from .env file
const apiKey = import.meta.env.VITE_MAPS_API;

//get gps data using store getter (reactive)
const gpsData = computed(() => gpsStore.getGPSData);

//load google maps and initialize map
const initializeMap = async () => {
  try {
    //set global google variable to maps object
    google = await loadGoogleMaps(apiKey);

    //give map default center of los angeles
    const center = { lat: 34.0522, lng: -118.2437 }; // Los Angeles
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
  if (!google || !map.value) return; //make sure maps is loaded

  gpsData.forEach((device) => {
    if (device.latest_device_point) {
      const { lat, lng } = device.latest_device_point;

      //update existing markers
      if (markers[device.device_id]) {
        markers[device.device_id].setPosition({ lat, lng });
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
                    <p>Status: ${device.online ? 'Online' : 'Offline'}</p>`,
        });

        //when marker clicked show info window
        marker.addListener('click', () => {
          infoWindow.open(map.value, marker);
        });

        //store marker in markers object
        markers[device.device_id] = marker;
      }
    }
  });
};

//when gps data changes update markers
watch(gpsData, (newData) => {
  updateMarkers(newData);
});

//initialize map on component mount
onMounted(async () => {
  console.log("API Key:", apiKey); //log the api key to make sure it's loaded
  await initializeMap(); //initialize google map
});
</script>