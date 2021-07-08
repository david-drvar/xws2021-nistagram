package setup

import (
	"github.com/david-drvar/xws2021-nistagram/chat_service/controllers"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/gorilla/mux"
	"net/http"
)

func ServerSetup(controller *controllers.MessageController) {
	go h.run()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from agent application!"))
	}).Methods("GET")

	router.HandleFunc("/ws/{roomId}", func(w http.ResponseWriter, r *http.Request) {
		roomId := mux.Vars(r)["roomId"]
		serveWs(w, r, roomId, controller.Service)
	})

	router.HandleFunc("/delete/{id}", controller.DeleteMessage).Methods("DELETE")
	router.HandleFunc("/room/{roomId}/messages", controller.GetMessagesForChatRoom).Methods("GET")
	router.HandleFunc("/room/{userId}", controller.GetChatRoomsForUser).Methods("GET")
	router.HandleFunc("/room", controller.CreateChatRoom).Methods("POST")
	router.HandleFunc("/room/conversation", controller.StartConversation).Methods("POST")
	router.HandleFunc("/request/accept", controller.AcceptMessageRequest).Methods("POST")
	router.HandleFunc("/request/decline", controller.DeclineMessageRequest).Methods("POST")
	router.HandleFunc("/message/{messageId}/seenPhoto", controller.SeenPhotoOrVideo).Methods("GET")

	c := common.SetupCors()

	http.Handle("/", c.Handler(router))
	http.ListenAndServe(":8003", c.Handler(router))

}
