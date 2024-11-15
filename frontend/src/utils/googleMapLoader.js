let isLoaded = false
let loadPromise

//load google maps api
export function loadGoogleMaps(apiKey) {
  //if script is already loaded, return the promise
  if (isLoaded) return loadPromise

  //if script is not loaded, create a new promise
  loadPromise = new Promise((resolve, reject) => {
    const script = document.createElement('script') //create a script element
    script.src = `https://maps.googleapis.com/maps/api/js?key=${apiKey}` //set the script source to google maps api
    script.async = true //set the script to async
    script.defer = true //set the script to defer

    //when the script is loaded, resolve the promise
    script.onload = () => {
      isLoaded = true
      resolve(window.google)
    }
    script.onerror = (error) => reject(error)
    document.head.appendChild(script)
  })

  return loadPromise
}
