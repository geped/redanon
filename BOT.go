pacchetto principale

importazione (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"net / http"
	"os"
	"stringhe"
)

const (
	MIN_WORDS_IN_TOKEN_FORM  =  2
	TOKEN_POSITION_IN_FORM   =  1
)

var  Bot  * tgbotapi. BotAPI

var  chatControlKeyboardUS  =  tgbotapi . NewReplyKeyboard (
	tgbotapi . NewKeyboardButtonRow (
		tgbotapi . NewKeyboardButton ( "Nuova chat" ),
		tgbotapi . NewKeyboardButton ( "Lascia chat / room" ),
	),
	tgbotapi . NewKeyboardButtonRow (
		tgbotapi . NewKeyboardButton ( "Crea una stanza segreta" ),
		tgbotapi . NewKeyboardButton ( "Entra in una stanza segreta" ),
	),
)

var  chatControlKeyboardRU  =  tgbotapi . NewReplyKeyboard (
	tgbotapi . NewKeyboardButtonRow (
		tgbotapi . NewKeyboardButton ( "Начать чат" ),
		tgbotapi . NewKeyboardButton ( "Покинуть чат / комнату" ),
	),
	tgbotapi . NewKeyboardButtonRow (
		tgbotapi . NewKeyboardButton ( "Создать секретную комнату" ),
		tgbotapi . NewKeyboardButton ( "Войти в секретную комнату" ),
	),
)

var  chatControlKeyboardMD  =  tgbotapi . NewReplyKeyboard (
	tgbotapi . NewKeyboardButtonRow (
		tgbotapi . NewKeyboardButton ( "Incepe conversatia" ),
		tgbotapi . NewKeyboardButton ( "Părăsește chatul / camera" ),
	),
	tgbotapi . NewKeyboardButtonRow (
		tgbotapi . NewKeyboardButton ( "Creați o cameră secretă" ),
		tgbotapi . NewKeyboardButton ( "Intră în camera secretă" ),
	),
)

func  init () {
	var  err  errore

	Bot , err  =  tgbotapi . NewBotAPI ( os . Getenv ( "BOT_TOKEN" ))
	if  err  ! =  nil {
		log . Panico ( err )
	}

	log . Printf ( "Autorizzato sull'account% s" , Bot . Self . UserName )
}

func  BotUpdateLoop () {
	u  : =  tgbotapi . Nuovo aggiornamento ( 0 )
	u . Timeout  =  60

	_ , err  : =  Bot . SetWebhook ( tgbotapi . NewWebhook ( os . Getenv ( "PUBLIC_URL" )   +  "/"  +  os . Getenv ( "BOT_TOKEN" )))
	if  err  ! =  nil {
		log . Fatale ( err )
	}

	aggiornamenti  : =  Bot . ListenForWebhook ( "/"  +  os . Getenv ( "BOT_TOKEN" ))
	vai su  http . ListenAndServe ( "0.0.0.0:"  +  os . Getenv ( "PORT" ), zero )

	// aggiornamenti, err = Bot.GetUpdatesChan (u)
	// if err! = nil {
	// log.Panic (err)
	//}

	for  update  : =  range  updates {

		se  aggiornamento . Messaggio  ==  nil {
			Continua
		}

		se  ! aggiornamento . Messaggio . IsCommand () {
			se  ! IsUserExist ( int64 ( update . Message . From . ID )) {
				AddNewUser ( int64 ( update . Message . From . ID ))
			}

			cambia  aggiornamento . Messaggio . Testo {
			case  "Nuova chat" , "Начать чат" , "Incepe conversatia" :
				NewChatButton ( aggiornamento )
				Continua
			caso  "Lascia chat / room" , "Покинуть чат / комнату" , "Părăsește chatul / camera" :
				LeaveChatButton ( aggiornamento )
				Continua
			case  "Crea una stanza segreta" , "Создать секретную комнату" , "Creați o cameră secretă" :
				CreateSecretRoom ( aggiornamento )
				Continua
			case  "Entra in una stanza segreta" , "Войти в секретную комнату" , "Intră în camera secretă" :
				JoinRoomTokenMessage ( aggiornamento )
				Continua
			}

			se  stringhe . Contiene ( update . Message . Text , "token" ) {
				formData  : =  stringhe . Campi ( aggiornamento . Messaggio . Testo )

				se  len ( formData ) <  MIN_WORDS_IN_TOKEN_FORM {
					InvalidTokenFormError ( aggiornamento )
					Continua
				}

				token  : =  formData [ TOKEN_POSITION_IN_FORM ]

				se  GetRoomAuthor ( token ) ==  0 {
					InvalidTokenError ( aggiornamento )
					Continua
				}

				JoinSecretRoom ( aggiornamento , token )
				Continua
			}

			se  ! CheckUserChattingStatus ( int64 ( update . Message . From . ID )) {
				Continua
			}

			SendMessageToAnotherUser ( aggiornamento )
			Continua
		}

		cambia  aggiornamento . Messaggio . Comando () {
		case  "start" :
			StartCommand ( aggiornamento )
		}
	}
}

func  NewChatButton ( update tgbotapi. Update ) {
	if  usersChecks ( update ) {
		ritorno
	}

	userID  : =  int64 ( update . Message . From . ID )
	ChangeUserSearchingStatus ( true , userID )

	msg  : =  tgbotapi . NewMessage ( userID , "Ricerca iniziato" )
	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg  =  tgbotapi . NewMessage ( userID , "Поиск начался" )
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg  =  tgbotapi . NewMessage ( userID , "cautarea un început" )
	}

	BotSendMessage ( msg )
}

func  LeaveChatButton ( update tgbotapi. Update ) {
	userID  : =  int64 ( update . Message . From . ID )

	se  ! IsUserExist ( userID ) {
		AddNewUser ( userID )
	}

	se  CheckIsRoomAuthor ( userID ) {
		DeleteRoom ( GetRoomToken ( userID ))
		DeleteRoomMessage ( aggiornamento , userID )
		ritorno
	}

	se  ! CheckUserChattingStatus ( userID ) {
		NotChattingError ( ID utente , aggiornamento )
		ritorno
	}

	RemoveChat ( aggiornamento , userID )
}

func  RemoveChat ( aggiornamento tgbotapi. Aggiornamento , userID  Int64 ) {
	secondUser  : =  GetSecondUser ( userID )

	DeleteChat ( userID , secondUser )

	ChangeUserChattingStatus ( falso , userID , secondUser )

	LeaveChatMessage ( aggiornamento , userID , secondUser )
}

func  CreateSecretRoom ( update tgbotapi. Update ) {
	if  usersChecks ( update ) {
		ritorno
	}

	token  : =  AddRoom ( int64 ( update . Message . From . ID ))

	CreateRoomMessage ( aggiornamento , token )
}

func  JoinSecretRoom ( update tgbotapi. Update , token  string ) {
	if  usersChecks ( update ) {
		ritorno
	}

	roomAuthor  : =  GetRoomAuthor ( token )
	DeleteRoom ( token )

	ChangeUserChattingStatus ( true , int64 ( update . Message . From . ID ), roomAuthor )

	AddChat ( int64 ( update . Message . From . ID ), roomAuthor )

	JoinRoomMessage ( aggiornamento , roomAuthor )
}

func  usersChecks ( update tgbotapi. Update ) bool {
	se  ! IsUserExist ( int64 ( update . Message . From . ID )) {
		AddNewUser ( int64 ( update . Message . From . ID ))
	}

	if  CheckUserSearchingStatus ( int64 ( update . Message . From . ID )) {
		AlreadySearchingError ( int64 ( update . Message . From . ID ), update )
		restituire  vero
	}

	if  CheckUserChattingStatus ( int64 ( update . Message . From . ID )) {
		AlreadyChattingError ( int64 ( update . Message . From . ID ), update )
		restituire  vero
	}

	if  CheckIsRoomAuthor ( int64 ( update . Message . From . ID )) {
		RoomAuthorError ( int64 ( update . Message . From . ID ), update )
		restituire  vero
	}
	return  false
}

func  StartCommand ( update tgbotapi. Update ) {
	se  ! IsUserExist ( int64 ( update . Message . From . ID )) {
		AddNewUser ( int64 ( update . Message . From . ID ))
	}

	msg  : =  tgbotapi . NewMessage ( int64 ( update . Message . From . ID ), "" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Text  =  "Привет, это Freenon чат -. Анонимный чат, где ты можешь высказывать свои мысли без последствий \ n \ n "  +
			"Чтобы начать чат с незнакомцем, нажми кнопку \" Начать чат \ " \ n \ n "  +
			"Чтобы покинуть чат, нажми кнопку \" Покинуть чат / комнату \ " \ n \ n "  +
			"Бот не сохраняет данные о пользователях, так что твои личные данные в безопасности. \ N \ n "  +
			"Если ты хочешь посмотреть, как работает бот - вот мое видео (https://www.youtube.com/watch?v=drtAdOByW54&t=1s) \ n \ n "  +
			"Если у тебя есть вопросы или предложения, пожалуйста, свяжись со мной, @YUART \ n \ n "
		msg . ReplyMarkup  =  chatControlKeyboardRU
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Bună, acesta este chatul AnonChatMoldova - un chat anonim în care îți poți exprima gândurile fără consecințe. \ N \ n "  +
			"Pentru a începe un chat cu un străin, do click pe butonul \" Incepe conversatia \ " . \ N \ n "  +
			"Pentru a părăsi chatul, dă click pe butonul \" Părăsește chatul / camera \ " . \ N \ n "  +
			"Botul nu salvează date despre utilizatori astfel încât informațiile dvs. personale să fie în siguranță. \ N \ n "  +
			"Dacă aveți întrebări sau sugestii, vă rugăm să mă contactați @OWNERMD"
		msg . ReplyMarkup  =  chatControlKeyboardMD
	} altro {
		msg . Text  =  "Ciao, questa è la chat di libertà, dove puoi esprimere liberamente la tua mente e parlare con altri sconosciuti. \ N \ n "  +
			"Per avviare la chat, premere \" Nuova chat "\ pulsante \ n \ n "  +
			"Per uscire dalla chat, premere \" Abbandona la chat / camera "\ pulsante \ n \ n "  +
			"Il bot non memorizza alcun dato personale, quindi le chat sono completamente anonime. \ N \ n "  +
			"Se vuoi controllare come funziona il bot, controlla il mio video (https://www.youtube.com/watch?v=drtAdOByW54&t=1s) \ n \ n "  +
			"Se hai qualche domanda o suggerimento, non esitare a contattarci, @YUART \ n \ n "
		msg . ReplyMarkup  =  chatControlKeyboardUS
	}

	BotSendMessage ( msg )
}

func  SendMessageToAnotherUser ( update tgbotapi. Update ) {
	secondUser  : =  GetSecondUser ( int64 ( update . Message . From . ID ))

	var  msg tgbotapi. Chattable

	se  aggiornamento . Messaggio . Testo  ! =  "" {
		msg  =  tgbotapi . NewMessage ( secondUser , update . Message . Text )
	}

	se  aggiornamento . Messaggio . Foto  ! =  Nil {
		foto  : =  tgbotapi . NewPhotoShare ( secondUser , "" )
		foto . FileID  = ( * aggiornamento . Messaggio . Foto ) [ 1 ]. FileID

		msg  =  foto
	}

	se  aggiornamento . Messaggio . Voce  ! =  Nil {
		voce  : =  tgbotapi . NewVoiceShare ( secondUser , "" )
		voce . FileID  =  aggiornamento . Messaggio . Voice . FileID

		msg  =  voce
	}

	se  aggiornamento . Messaggio . Animazione  ! =  Nil {
		gif  : =  tgbotapi . NewAnimationShare ( secondUser , "" )
		gif . FileID  =  aggiornamento . Messaggio . Animazione . FileID

		msg  =  gif
	}

	se  aggiornamento . Messaggio . Audio  ! =  Nil {
		audio  : =  tgbotapi . NewAudioShare ( secondUser , "" )
		audio . FileID  =  aggiornamento . Messaggio . Audio . FileID

		msg  =  audio
	}

	se  aggiornamento . Messaggio . Sticker  ! =  Nil {
		adesivo  : =  tgbotapi . NewStickerShare ( secondUser , "" )
		adesivo . FileID  =  aggiornamento . Messaggio . Adesivo . FileID

		msg  =  adesivo
	}

	se  aggiornamento . Messaggio . Documento  ! =  Nil {
		doc  : =  tgbotapi . NewDocumentShare ( secondUser , "" )
		doc . FileID  =  aggiornamento . Messaggio . Documento . FileID

		msg  =  doc
	}

	se  aggiornamento . Messaggio . Video  ! =  Nil {
		video  : =  tgbotapi . NewVideoShare ( secondUser , "" )
		video . FileID  =  aggiornamento . Messaggio . Video . FileID

		msg  =  video
	}

	se  aggiornamento . Messaggio . VideoNote  ! =  Nil {
		videoNote  : =  tgbotapi . NewVideoNoteShare ( secondUser , 0 , "" )
		videoNote . FileID  =  aggiornamento . Messaggio . VideoNote . FileID
		videoNote . Lunghezza  =  aggiornamento . Messaggio . VideoNote . Lunghezza

		msg  =  videoNote
	}

	se  msg  ==  nil {
		se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
			msg  =  tgbotapi . NewMessage ( int64 ( update . Message . From . ID ), "Бот не может это отправить. Пожалуйста, свяжитесь с администрацией" )
		} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
			msg  =  tgbotapi . NewMessage ( int64 ( aggiornamento . Messaggio . Da . ID ), "Botul nu poate trimite asta. Vă rugăm să contactați administrația" )
		} altro {
			msg  =  tgbotapi . NewMessage ( int64 ( update . Message . From . ID ), "Il bot non può ancora inviarlo! Per favore, contatta il creatore" )
		}
	}

	BotSendMessage ( msg )
}

func  BotSendMessage ( message tgbotapi. Chattable ) {
	var  err  errore

	_ , err  =  Bot . Invia ( messaggio )

	if  err  ! =  nil {
		BotSendMessageError ( err )
	}
}

func  JoinRoomTokenMessage ( update tgbotapi. Update ) {
	if  usersChecks ( update ) {
		ritorno
	}

	msg  : =  tgbotapi . NewMessage ( int64 ( update . Message . From . ID ), "Fornisci ora il token segreto. Digita la parola" token ", spazio e fornisci token" )
	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Text  =  "Теперь введите секретный токен. Пожалуйста, напишите слово 'token', пробел, и введите токен"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Acum introduceți simbolul secret. Vă rugăm să scrieți cuvântul" token ", un spațiu și introduceți simbolul"
	}

		BotSendMessage ( msg )
}

func  LeaveChatMessage ( aggiornamento tgbotapi. Aggiornamento , userID  Int64 , secondUser  Int64 ) {
	msg  : =  tgbotapi . NewMessage ( userID , "Tu leaved una chat" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg  =  tgbotapi . NewMessage ( userID , "Вы покинули чат" )
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg  =  tgbotapi . NewMessage ( userID , "Ai Parasit chatul" )
	}

	BotSendMessage ( msg )
	BotSendMessage ( tgbotapi . NewMessage ( secondUser , "Lo sconosciuto lascia la chat" ))
}

func  DeleteRoomMessage ( aggiornamento tgbotapi. Aggiornamento , userID  Int64 ) {
	msg  : =  tgbotapi . NewMessage ( userID , "si elimina una stanza segreta" )
	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg  =  tgbotapi . NewMessage ( userID , "Вы удалили секретную комнату" )
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg  =  tgbotapi . NewMessage ( userID , "ATI steri secreta fotocamera" )
	}

	BotSendMessage ( msg )
}

func  CreateRoomMessage ( update tgbotapi. Update , token  string ) {
	msg  : =  tgbotapi . NewMessage ( int64 ( update . Message . From . ID ), fmt . Sprintf ( "Hai creato una stanza segreta. Il tuo token segreto% v. Condividilo con un'altra persona." , Token ))

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Testo  =  fmt . Sprintf ( "Вы создали секретную комнату. Ваш секретный токен% v. Передайте его другому человеку, что бы он мог подключиться к комнате." , Token )
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Testo  =  fmt . Sprintf ( "Ați creat o cameră secretă. Token-ul tău secret% v. Dă-i altei persoane, astfel încât să se poată conecta la cameră." , Token )
	}

	BotSendMessage ( msg )
}

func  JoinRoomMessage ( update tgbotapi. Update , roomAuthor  int64 ) {
	msg  : =  tgbotapi . NewMessage ( int64 ( aggiornamento . Messaggio . Da . ID ), "Sei entrato in una stanza segreta" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Text  =  "Вы присоединились к секретной комнате"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Te-ai alăturat camerei secrete"
	}

	BotSendMessage ( msg )
	BotSendMessage ( tgbotapi . NewMessage ( roomAuthor , "Un'altra persona si è unita a una stanza segreta" ))
}

func  BotSendMessageError ( errore err  ) {
	fmt . Println ( fmt . Errorf ( "Invio messaggio non riuscito:% w \ n " , err ))
}

func  InvalidTokenError ( update tgbotapi. Update ) {
	msg  : =  tgbotapi . NewMessage ( int64 ( update . Message . From . ID ), "Impossibile trovare una stanza segreta con quel token" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Testo  =  "Не могу найти комнату с таким токеном"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Nu găsesc o cameră cu un astfel de Token"
	}

	BotSendMessage ( msg )
}

func  InvalidTokenFormError ( update tgbotapi. Update ) {
	msg  : =  tgbotapi . NewMessage ( int64 ( aggiornamento . Messaggio . Da . ID ), "Il token è stato fornito in modo errato" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Testo  =  "Токен был введен не правильно"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Token-ul a fost introdus incorect"
	}

	BotSendMessage ( msg )
}

func  AlreadySearchingError ( utente  int64 , aggiorna tgbotapi. Aggiorna ) {
	msg  : =  tgbotapi . NewMessage ( utente , "Stai già cercando" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Testo  =  "Вы уже ищите чат"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Căutați deja un chat"
	}

	BotSendMessage ( msg )
}

func  AlreadyChattingError ( utente  int64 , aggiorna tgbotapi. Aggiorna ) {
	msg  : =  tgbotapi . NewMessage ( utente , "Stai già chattando" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Testo  =  "Вы уже в чате"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Sunteti deja in chat"
	}

	BotSendMessage ( msg )
}

func  RoomAuthorError ( user  int64 , update tgbotapi. Update ) {
	msg  : =  tgbotapi . NewMessage ( utente , "Sei l'autore della stanza segreta. Attendi una seconda persona o esci dalla stanza" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Text  =  "Вы автор секретной комнаты. Подождите второго человека или выйдите из комнаты"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Ești autorul camerei secrete. Așteptați a doua persoană sau părăsiți camera"
	}

	BotSendMessage ( msg )
}

func  NotChattingError ( user  int64 , update tgbotapi. Update ) {
	msg  : =  tgbotapi . NewMessage ( utente , "Non stai chattando adesso!" )

	se  aggiornamento . Messaggio . Da . LanguageCode  ==  "ru"  ||  aggiornamento . Messaggio . Da . LanguageCode  ==  "ua" {
		msg . Testo  =  "Вы не находитесь в чате"
	} altrimenti  se  aggiorna . Messaggio . Da . LanguageCode  ==  "ro" {
		msg . Text  =  "Nu va aflati in niciun chat"
	}

	BotSendMessage ( msg )
}
