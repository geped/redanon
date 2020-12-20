pacchetto principale

importazione (
	"crypto / md5"
	"codifica / esadecimale"
	"fmt"
	"tempo"
)

type  UserStatuses  struct {
	SearchingStatus  bool
	ChattingStatus   bool
}

func ( user  * UserStatuses ) SetSearchingStatus ( state  bool ) {
	utente . SearchingStatus  =  state
}

func ( user  * UserStatuses ) SetChattingStatus ( state  bool ) {
	utente . ChattingStatus  =  stato
}

func ( user  * UserStatuses ) IsUserSearching () bool {
	 utente di ritorno . SearchingStatus
}

func ( user  * UserStatuses ) IsUserChatting () bool {
	 utente di ritorno . ChattingStatus
}

var  Mappa utenti  [ int64 ] * UserStatuses
var  Chats  mappa [ int64 ] int64
var  Rooms  mappa [ stringa ] int64

func  init () {
	Utenti  =  make ( map [ int64 ] * UserStatuses )
	per  chiave , val  : =  range  GetUsersFromDB () {
		Utenti [ chiave ] =  val
	}

	Chat  =  make ( map [ int64 ] int64 )
	per  chiave , val  : =  range  GetChatsFromDB () {
		Chat [ chiave ] =  val
	}

	Stanze  =  make ( map [ string ] int64 )
	per  chiave , val  : =  range  GetRoomsFromDB () {
		Camere [ chiave ] =  val
	}
}

func  GetUsersCache () map [ int64 ] * UserStatuses {
	restituire gli  utenti
}

func  GetChatsCache () map [ int64 ] int64 {
	restituire le  chat
}

func  GetRoomsCache () map [ stringa ] int64 {
	ritorno  Camere
}

func  IsUserExist ( user  int64 ) bool {
	_ , esistono  : =  Users [ utente ]
	il ritorno  esiste
}

func  AddNewUser ( user  int64 ) {
	Utenti [ utente ] =  nuovo ( UserStatuses )
	Utenti [ utente ]. SetSearchingStatus ( false )
	Utenti [ utente ]. SetChattingStatus ( false )
}

func  ChangeUserSearchingStatus ( status  bool , users  ... int64 ) {
	per  _ , user  : =  range  users {
		Utenti [ utente ]. SetSearchingStatus ( stato )
	}
}

func  ChangeUserChattingStatus ( status  bool , users  ... int64 ) {
	per  _ , user  : =  range  users {
		Utenti [ utente ]. SetChattingStatus ( stato )
	}
}

func  CheckUserSearchingStatus ( user  int64 ) bool {
	return  Users [ utente ]. IsUserSearching ()
}

func  CheckUserChattingStatus ( user  int64 ) bool {
	return  Users [ utente ]. IsUserChatting ()
}

func  SearchingUsersList () ( users [] int64 ) {
	per  userId , userStatus  : =  range  Users {
		se  userStatus . IsUserSearching () {
			utenti  =  append ( utenti , userId )
		}
	}

	ritorno
}

func  AddChat ( firstUser , secondUser  int64 ) {
	Chat [ firstUser ] =  secondUser
	Chat [ secondUser ] =  firstUser
}

func  DeleteChat ( firstUser , secondUser  int64 ) {
	elimina ( chat , firstUser )
	elimina ( chat , secondUser )
}

func  GetSecondUser ( firstUser  int64 ) int64 {
	return  Chats [ firstUser ]
}

func  AddRoom ( authorUser  int64 ) ( token  string ) {
	token  =  CreateToken ( authorUser )
	Stanze [ token ] =  authorUser
	ritorno
}

func  DeleteRoom ( token  string ) {
	elimina ( stanze , token )
}

func  GetRoomAuthor ( token  string ) int64 {
	Return  Rooms [ token ]
}

func  GetRoomToken ( user  int64 ) stringa {
	for  token , roomAuthor  : =  range  Rooms {
		se  roomAuthor  ==  user {
			 token di ritorno
		}
	}

	ritorno  ""
}

func  CheckIsRoomAuthor ( user  int64 ) bool {
	per  _ , roomAuthor  : =  range  Rooms {
		se  roomAuthor  ==  user {
			restituire  vero
		}
	}

	return  false
}

func  CreateToken ( id  int64 ) stringa {
	hasher  : =  md5 . Nuovo ()
	hasher . Scrivi ([] byte ( fmt . Sprintf ( "% v% v" , id , time . Now ())))
	return  hex . EncodeToString ( hasher . Sum ( nil ))
}
