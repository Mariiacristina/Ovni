package schema


//Schema del marciano! parte en Mayuscula pq es public
type Marciano struct {
    Id int   `json:"id,omitempty"`
    Nombre string `json:"nombre,omitempty"`
}

type Nave struct{
  Id int   `json:"id,omitempty"`
  Nombre string `json:"nombre,omitempty"`
}

type Aero struct {
  Id int `json:"id,omitempty"`
  Nombre string `json:"nombre,omitempty"`
  Nave_o string `json:"nave_o,omitempty"`
  Nave_d string `json:"nave_d,omitempty"`
  Cant_max int `json:"cant_max,omitempty"`
}

type Viaje struct {
  Id_Pasajero int `json:"id_pasajero,omitempty"`
  Id_Aero int `json:"id_aero,omitempty"`
  Fecha string `json:"fecha,omitempty"`
}
