pacchetto principale

importazione (
	"contesto"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

var  DBConnection  =  func () ( connection  * pgx. Conn ) {
	var  err  errore

	connessione , err  =  pgx . Connetti ( context . Background (), os . Getenv ( "DB_TOKEN" ))
	if  err  ! =  nil {
		DBConnectionError ( err )
	}

	log . Println ( "Connesso a PSQL!" )

	ritorno
} ()

func  BackupData ( mappa utenti  [ int64 ] * UserStatuses , mappa chat [ int64 ] int64 , mappa stanze [ stringa ] int64 ) {  
	InsertUsersCache ( utenti )
	InsertChatsCache ( chat )
	InsertRoomsCache ( stanze )
}

func  InsertUsersCache ( mappa utenti  [ int64 ] * UserStatuses ) {
	per  chiave , val  : =  range  users {
		_ , err  : =  DBConnection . Exec ( context . Background (), "INSERT INTO users VALUES ($ 1, $ 2, $ 3)" ,
			chiave , val . IsUserSearching (), val . IsUserChatting ())

		if  err  ! =  nil {
			BackupCacheError ( utenti )
		}
	}
}

func  InsertChatsCache ( chat  map [ int64 ] int64 ) {
	_ , err  : =  DBConnection . Exec ( context . Background (), "DELETE FROM chats" )

	for  key , val  : =  range  chats {
		_ , err  =  DBConnection . Exec ( context . Background (), "INSERT INTO chats VALUES ($ 1, $ 2)" ,
			chiave , val )

		if  err  ! =  nil {
			DBQueryError ( err )
		}
	}
}

func  InsertRoomsCache ( room  map [ string ] int64 ) {
	_ , err  : =  DBConnection . Exec ( context . Background (), "DELETE FROM rooms" )

	for  key , val  : =  range  rooms {
		_ , err  =  DBConnection . Exec ( context . Background (), "INSERT INTO rooms VALUES ($ 1, $ 2)" ,
			chiave , val )

		if  err  ! =  nil {
			DBQueryError ( err )
		}
	}
}

func  BackupCacheError ( interfaccia cache  {}) {
	switch  v  : =  cache . ( type ) {
	case  map [ int64 ] * UserStatuses :
		for  key , val  : =  range  v {
			_ , err  : =  DBConnection . Exec ( context . Background (), "UPDATE users SET is_searching = $ 1, is_chatting = $ 2 WHERE user_id = $ 3" ,
				val . IsUserSearching (), val . IsUserChatting (), chiave )

			if  err  ! =  nil {
				DBQueryError ( err )
			}
		}
	}
}

func  GetUsersFromDB () ( mappa utenti  [ int64 ] * UserStatuses ) {
	utenti  =  make ( map [ int64 ] * UserStatuses )

	righe , err  : =  DBConnection . Query ( context . Background (), "SELECT * FROM users" )
	if  err  ! =  nil {
		DBQueryError ( err )
	}

	rinviare  righe . Chiudi ()

	per le  righe . Avanti () {
		var  n  int64
		 ricerca var , chat  bool

		err  =  righe . Scan ( & n , e la ricerca , e la chat )
		if  err  ! =  nil {
			DBQueryError ( err )
		}

		utenti [ n ] =  nuovo ( UserStatuses )
		utenti [ n ]. SetSearchingStatus ( ricerca )
		utenti [ n ]. SetChattingStatus ( chat )
	}

	ritorno
}

func  GetChatsFromDB () ( mappa chat  [ int64 ] int64 ) {
	chat  =  make ( map [ int64 ] int64 )

	righe , err  : =  DBConnection . Query ( context . Background (), "SELECT * FROM chats" )
	if  err  ! =  nil {
		DBQueryError ( err )
	}

	rinviare  righe . Chiudi ()

	per le  righe . Avanti () {
		var  firstUser , secondUser  int64

		err  =  righe . Scan ( e firstUser , e secondUser )
		if  err  ! =  nil {
			DBQueryError ( err )
		}

		chat [ firstUser ] =  secondUser
	}

	ritorno
}

func  GetRoomsFromDB () ( room  map [ string ] int64 ) {
	rooms  =  make ( map [ string ] int64 )

	righe , err  : =  DBConnection . Query ( context . Background (), "SELECT * FROM rooms" )
	if  err  ! =  nil {
		DBQueryError ( err )
	}

	rinviare  righe . Chiudi ()

	per le  righe . Avanti () {
		var  authorUser  int64
		 stringa token  var

		err  =  righe . Scansione ( & token , & authorUser )
		if  err  ! =  nil {
			DBQueryError ( err )
		}

		rooms [ token ] =  authorUser
	}

	ritorno
}

func  DBQueryError ( errore err  ) {
	log . Println ( fmt . Errorf ( "QueryRow non riuscita:% v \ n " , err ))
}

func  DBConnectionError ( errore err  ) {
	log . Println ( fmt . Errorf ( "Impossibile connettersi al database:% w \ n " , err ))
	panico ( err )
}
