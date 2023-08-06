package mock

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -destination=mock_usecase.go -package=mock github.com/hengkyawijaya/simple-go/usecase HelloUsecase,AuthUsecase
//go:generate mockgen -destination=mock_repository.go -package=mock github.com/hengkyawijaya/simple-go/repository ConfigRepository,IPInfoRepository
