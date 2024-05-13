package main

type AplicationFunctions interface{
	Open()
	Close()
}

type Aplication struct{
	app AplicationFunctions
}

func(apps Aplication)Run(){
	apps.app.Open()
	apps.app.Close()
}

func NewApplication(apps AplicationFunctions) *Aplication{
      return &Aplication{app:apps}
}

func main(){
	
}