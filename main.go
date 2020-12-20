pacchetto principale

importazione (
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"matematica / rand"
	"tempo"
)

const (
	PAUSE_FOR_BACKUPS           =  100
	PAUSE_FOR_CHATMAKER         =  1
	USERS_AMOUNT_FOR_CHATMAKER  =  2
)

func  main () {
	vai a  BackupCache ()
	vai a  ChatMaker ()
	BotUpdateLoop ()
}

func  BackupCache () {
	per {
		Pausa ( PAUSE_FOR_BACKUPS )

		log . Println ( "Backup" )
		BackupData ( utenti , chat , stanze virtuali )
	}
}

func  ChatMaker () {
	per  vero {
		freeUsers  : =  SearchingUsersList ()
		usersAmount  : =  len ( freeUsers )

		if  usersAmount  > =  USERS_AMOUNT_FOR_CHATMAKER {
			CreateNewChat ( usersAmount , freeUsers )
		}

		Pausa ( PAUSE_FOR_CHATMAKER )
	}
}

func  CreateNewChat ( usersAmount  int , freeUsers [] int64 ) {
	firstUser , secondUser  : =  ChooseRandomUsers ( usersAmount , freeUsers )

	MakeChat ( firstUser , secondUser )

	per  _ , utente  : =  range [] int64 { firstUser , secondUser } {
		BotSendMessage ( tgbotapi . NewMessage ( user , "Now you can chat" ))
	}
}

func  ChooseRandomUsers ( usersAmount  int , freeUsers [] int64 ) ( firstUser , secondUser  int64 ) {
	rand . Seed ( time . Now (). UnixNano ())
	userOne  : =  rand . Intn ( usersAmount )
	userTwo  : =  rand . Intn ( usersAmount )

	per  userTwo  ==  userOne {
		userTwo  =  rand . Intn ( usersAmount )
	}

	firstUser  =  freeUsers [ userOne ]
	secondUser  =  freeUsers [ userTwo ]
	ritorno
}

func  MakeChat ( firstUser  int64 , secondUser  int64 ) {
	ChangeUserSearchingStatus ( false , firstUser , secondUser )
	ChangeUserChattingStatus ( true , firstUser , secondUser )

	AddChat ( firstUser , secondUser )
}

func  Pausa ( secondi  uint16 ) {
	amt  : =  time . Durata ( secondi )
	tempo . Sleep ( time . Second  *  amt )
}
