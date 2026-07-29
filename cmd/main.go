package main
import "log"
import "os"

func main(){
	cfg := config{
		addr: ":8080",
		
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil{
		log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}

}