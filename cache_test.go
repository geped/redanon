pacchetto principale

importazione (
	"riflettere"
	"ordinare"
	"test"
)

func  Add5NewUsers () {
	per  i  : =  0 ; i  <  5 ; i ++ {
		AddNewUser ( int64 ( i ))
	}
}

func  Add5NewChats () {
	AddChat ( 1 , 2 )
	AddChat ( 3 , 4 )
	AddChat ( 5 , 6 )
	AddChat ( 7 , 8 )
	AddChat ( 9 , 10 )
}

funz  Add5Rooms () ( token [] stringa ) {
	per  i  : =  0 ; i  <  5 ; i ++ {
		token  =  append ( token , AddRoom ( int64 ( i )))
	}
	ritorno
}

func  TestAdd5NewUsers ( t  * testing. T ) {
	Add5NewUsers ()

	se  len ( GetUsersCache ()) ! =  5 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 5 , len ( GetUsersCache ()))
	}

	per  _ , userStates  : =  range  GetUsersCache () {
		se  userStates . IsUserSearching () ||  userStates . IsUserChatting () {
			t . Errorf ( "Previsto% v, ottenuto% v" , falso , vero )
		}
	}
}

func  TestAdd5NewUsersIs7UsersExist ( t  * testing. T ) {
	Add5NewUsers ()

	var  risponde [] bool
	per  i  : =  0 ; i  <  7 ; i ++ {
		se  IsUserExist ( int64 ( i )) {
			risposte  =  append ( risposte , IsUserExist ( int64 ( i )))
		}
	}

	if  len ( risposte ) ! =  5 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 5 , len ( risposte ))
	}
}

func  TestAdd5NewUsers3UsersChangeSearchingStatus ( t  * testing. T ) {
	Add5NewUsers ()

	per  i  : =  0 ; i  <  3 ; i ++ {
		ChangeUserSearchingStatus ( int64 ( i ), true )
	}

	var  risponde [] bool
	per  _ , stati  : =  range  GetUsersCache () {
		se gli  stati . IsUserSearching () {
			risposte  =  aggiungi ( risposte , stati . IsUserSearching ())
		}
	}

	if  len ( risposte ) ! =  3 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 3 , len ( risposte ))
	}
}

func  TestAdd5NewUsers3UsersChangeChattingStatus ( t  * testing. T ) {
	Add5NewUsers ()

	per  i  : =  0 ; i  <  3 ; i ++ {
		ChangeUserChattingStatus ( int64 ( i ), true )
	}

	var  risponde [] bool
	per  _ , stati  : =  range  GetUsersCache () {
		se gli  stati . IsUserChatting () {
			risposte  =  aggiungi ( risposte , stati . IsUserChatting ())
		}
	}

	if  len ( risposte ) ! =  3 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 3 , len ( risposte ))
	}
}

func  TestAdd5NewUsers3UsersChangeSearchingStatusCheckSearchingStatus ( t  * testing. T ) {
	Add5NewUsers ()

	per  i  : =  0 ; i  <  3 ; i ++ {
		ChangeUserSearchingStatus ( int64 ( i ), true )
	}

	var  risponde [] bool
	per  i  : =  0 ; i  <  5 ; i ++ {
		se  CheckUserSearchingStatus ( int64 ( i )) {
			risposte  =  append ( risposte , CheckUserSearchingStatus ( int64 ( i )))
		}
	}

	if  len ( risposte ) ! =  3 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 3 , len ( risposte ))
	}
}

func  TestAdd5NewUsers3UsersChangeChattingStatusCheckChattingStatus ( t  * testing. T ) {
	Add5NewUsers ()

	per  i  : =  0 ; i  <  3 ; i ++ {
		ChangeUserChattingStatus ( int64 ( i ), true )
	}

	var  risponde [] bool
	per  i  : =  0 ; i  <  5 ; i ++ {
		se  CheckUserChattingStatus ( int64 ( i )) {
			risposte  =  append ( risposte , CheckUserChattingStatus ( int64 ( i )))
		}
	}

	if  len ( risposte ) ! =  3 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 3 , len ( risposte ))
	}
}

func  TestAdd5NewUsers3UsersChangeSearchingStatusGetSearchingUsersList ( t  * testing. T ) {
	Add5NewUsers ()

	per  i  : =  0 ; i  <  3 ; i ++ {
		ChangeUserSearchingStatus ( int64 ( i ), true )
	}

	risposte  : = [] int { 0 , 1 , 2 }
	utenti  : =  SearchingUsersList ()

	var  iusers [] int
	per  _ , i  : =  range  users {
		iusers  =  append ( iusers , int ( i ))
	}

	ordina . Ints ( iusers )
	se  ! riflettere . DeepEqual ( risposte , iusers ) {
		t . Errorf ( "Previsto% v, ottenuto% v" , risposte , iusers )
	}
}

func  TestAdd5Chats ( t  * testing. T ) {
	Add5NewChats ()

	se  len ( GetChatsCache ()) ! =  10 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 10 , len ( GetChatsCache ()))
	}

}

func  TestAdd5ChatsDelete3Chats ( t  * testing. T ) {
	Add5NewChats ()

	DeleteChat ( 1 , 2 )
	DeleteChat ( 5 , 6 )
	DeleteChat ( 9 , 10 )

	se  len ( GetChatsCache ()) ! =  4 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 10 , len ( GetChatsCache ()))
	}

}

func  TestAdd5Chats3TimesGetSecondUser ( t  * testing. T ) {
	Add5NewChats ()

	var  users [] int64

	utenti  =  append ( utenti , GetSecondUser ( 1 ))
	utenti  =  append ( utenti , GetSecondUser ( 10 ))
	utenti  =  append ( utenti , GetSecondUser ( 5 ))

	risposta  : = [] int64 { 2 , 9 , 6 }

	se  ! riflettere . DeepEqual ( risposta , utenti ) {
		t . Errorf ( "Previsto% v, ottenuto% v" , risposta , utenti )
	}
}

func  TestAdd5NewRooms ( t  * testing. T ) {
	Add5Rooms ()

	se  len ( GetRoomsCache ()) ! =  5 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 5 , len ( GetRoomsCache ()))
	}
}

func  TestAdd5NewRoomsDelete3Rooms ( t  * testing. T ) {
	gettoni  : =  Add5Rooms ()

	DeleteRoom ( token [ 0 ])
	DeleteRoom ( token [ 1 ])
	DeleteRoom ( token [ 2 ])

	se  len ( GetRoomsCache ()) ! =  2 {
		t . Errorf ( "Previsto% v, ottenuto% v" , 2 , len ( GetRoomsCache ()))
	}
}

func  TestAdd5NewRoomsGet3RoomsAuthor ( t  * testing. T ) {
	gettoni  : =  Add5Rooms ()

	var  users [] int64
	utenti  =  append ( utenti , GetRoomAuthor ( token [ 0 ]))
	utenti  =  append ( utenti , GetRoomAuthor ( token [ 3 ]))
	utenti  =  append ( utenti , GetRoomAuthor ( token [ 1 ]))

	risposta  : = [] int64 { 0 , 3 , 1 }

	se  ! riflettere . DeepEqual ( risposta , utenti ) {
		t . Errorf ( "Previsto% v, ottenuto% v" , risposta , utenti )
	}
}

func  TestAdd5NewRoomsGet3FirstTokens ( t  * testing. T ) {
	gettoni  : =  Add5Rooms ()

	usersAnsw  : = [] int64 { 0 , 1 , 2 }
	var  userTokens [] stringa

	per  _ , id  : =  range  usersAnsw {
		userTokens  =  append ( userTokens , GetRoomToken ( id ))
	}

	risposta  : = [] stringa { token [ 0 ], token [ 1 ], token [ 2 ]}

	se  ! riflettere . DeepEqual ( answer , userTokens ) {
		t . Errorf ( "Previsto% v, ottenuto% v" , risposta , userTokens )
	}
}

func  TestAdd5NewRoomsCheck3UsersIsRoomAuthor ( t  * testing. T ) {
	Add5Rooms ()

	usersAnsw  : = [] int64 { 0 , 1 , 6 }
	var  userAuthor [] bool

	per  _ , id  : =  range  usersAnsw {
		userAuthor  =  append ( userAuthor , CheckIsRoomAuthor ( id ))
	}

	risposta  : = [] bool { true , true , false }

	se  ! riflettere . DeepEqual ( answer , userAuthor ) {
		t . Errorf ( "Previsto% v, ottenuto% v" , risposta , userAuthor )
	}
}
