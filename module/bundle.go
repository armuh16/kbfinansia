package module

import (
	//Route
	authRoute "github.com/armuh16/kbfinansia/module/auth/route"
	transactionRoute "github.com/armuh16/kbfinansia/module/transaction/route"
	userRoute "github.com/armuh16/kbfinansia/module/user/route"

	//Logic
	authLogic "github.com/armuh16/kbfinansia/module/auth/logic"
	transactionLogic "github.com/armuh16/kbfinansia/module/transaction/logic"
	userLogic "github.com/armuh16/kbfinansia/module/user/logic"

	//Repository
	transactionRepository "github.com/armuh16/kbfinansia/module/transaction/repository"
	userRepository "github.com/armuh16/kbfinansia/module/user/repository"

	"go.uber.org/fx"
)

// Register Route
var BundleRoute = fx.Options(
	fx.Invoke(transactionRoute.NewRoute),
	fx.Invoke(userRoute.NewRoute),
	fx.Invoke(authRoute.NewRoute),
)

// Register logic
var BundleLogic = fx.Options(
	fx.Provide(userLogic.NewLogic),
	fx.Provide(transactionLogic.NewLogic),
	fx.Provide(authLogic.NewLogic),
)

// Register Repository
var BundleRepository = fx.Options(
	fx.Provide(userRepository.NewRepository),
	fx.Provide(transactionRepository.NewRepository),
)
