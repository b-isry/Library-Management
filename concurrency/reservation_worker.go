package concurrency

import (
	"Library_Management/models"
	"Library_Management/services"
	"fmt"
	"sync"
)

type ReservationHandler struct {
	library          *services.Library
	reservationQueue chan models.Reservation
	workerCount      int
	wg               sync.WaitGroup
}

func NewReservationHandler(library *services.Library, workerCount int) *ReservationHandler {
	return &ReservationHandler{
		library:          library,
		reservationQueue: make(chan models.Reservation, 100),
		workerCount:      workerCount,
	}
}

func (h *ReservationHandler) Start() {
	for i := 0; i < h.workerCount; i++ {
		h.wg.Add(1)
		go h.worker()
	}
}

func (h *ReservationHandler) Stop() {
	close(h.reservationQueue)
	h.wg.Wait()
}

func (h *ReservationHandler) SubmitReservation(reservation models.Reservation) {
	h.reservationQueue <- reservation
}

func (h *ReservationHandler) worker() {
	defer h.wg.Done()

	for reservation := range h.reservationQueue {
		err := h.library.ReserveBook(reservation.MemberID, reservation.BookID)
		if err != nil {
			fmt.Printf("Failed to process reservation: %v\n", err)
		}
	}
}
