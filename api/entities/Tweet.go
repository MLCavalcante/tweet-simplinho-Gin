package entities

import (
	"github.com/pborman/uuid"

)	//! Tweet maiúsculo = struct é publica para ser importada pelo controller 

type Tweet struct {
	ID string `json:"id"`  // crases para renomear quando for json queremos id minúsculo
	Description string `json:"description"`
}

// funcção que vai retornar a nova struct de tweet

func NewTweet() *Tweet { // vai retornar um ponteiro de tweet
	tweet := Tweet{  // instanciar uma var vou dizer que ela é do tipo tweet
      ID: uuid.New(), //aqui vamos iniciar com um id passando uuid
	} 
	return &tweet  
}