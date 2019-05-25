package main

import "net/http"
import "encoding/json"
import "strings"

type weatherData struct{
  Name string `json:"name"`
  Main struct{
    Kelvin float64 `json:"temp"`
  } `json:"main"`

}

func main(){
   http.HandleFunc("/hi", hi)
   http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
   city := strings.SplitN(r.URL.Path, "/", 3)[2]

   data, err := query(city)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
   }

   w.Header().Set("Content-Type", "application/json; charset=utf-8")
   json.NewEncoder(w).Encode(data)
})






http.HandleFunc("/underground-weather/", func(w http.ResponseWriter, r *http.Request) {
coordinates:=strings.SplitN(r.URL.Path, "/", 3)[2]
data, err := query1(coordinates)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

w.Header().Set("Content-Type", "application/json; charset=utf-8")
json.NewEncoder(w).Encode(data)
})



   http.HandleFunc("/", hello)
   http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("hello"))
}

func hi(w http.ResponseWriter, r *http.Request){
  // ???
  w.Write([]byte("hi Andela"))

}
func query(city string) (weatherData, error) {
   resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=e0fc2d0b038108cd17f1b71026ece2ce&q=" + city)
   if err != nil {
       return weatherData{}, err
   }

   defer resp.Body.Close()
   var d weatherData

   if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
       return weatherData{}, err
   }

   return d, nil
}




func query1(coordinates string) (weatherData, error) {
   resp, err := http.Get("https://api.darksky.net/forecast/f0675ee17aaac952208fdb9c0d3854d1/"+coordinates)
   if err != nil {
       return weatherData{}, err
   }

   defer resp.Body.Close()
   var d weatherData

   if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
       return weatherData{}, err
   }

   return d, nil
}
