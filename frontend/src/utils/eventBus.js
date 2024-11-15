import mitt from 'mitt';

//decided to utilize this event emitter to handle the communication between components. I tried to use the 'watch' vue function to look for changes in the device visibility to retrigger the updateMarker function in GoogleMap.vue but it would only work when making a device visible and not invisible. this allowed for more simplicity in the event handling 
const eventBus = mitt();
export default eventBus;
