package starter

import (
	"context"
	"geulSsi/app/repository"
	"geulSsi/app/service"
	"geulSsi/config/db"
)

func WiseSayingStarter() {
	ctx := context.TODO()
	wiseSayingService := wiseSayingServiceDI()
	wiseSayingService.RegisterWiseSayingChannel(ctx)

}

func wiseSayingServiceDI() service.WiseSayingService {
	return service.NewWiseSayingService(
		repository.NewUserRepository(*db.MySQL),
		repository.NewWiseSayingRepository(*db.MySQL),
	)
}
