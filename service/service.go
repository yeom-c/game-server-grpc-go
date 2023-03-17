package service

import (
	account_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/account"
	asset_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/asset"
	battle_result_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/battle_result"
	battle_user_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/battle_user"
	character_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/character"
	character_broadcast_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/character_broadcast"
	character_collection_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/character_collection"
	cheat_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/cheat"
	costume_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/costume"
	deck_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/deck"
	fate_card_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/fate_card"
	gacha_log_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/gacha_log"
	item_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/item"
	mail_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/mail"
	recipe_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/recipe"
	shop_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/shop"
	story_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/story"
	user_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/user"

	"github.com/yeom-c/game-server-grpc-go/config"
	"google.golang.org/grpc"
)

func NewService(grpcServer *grpc.Server) {
	account_pb.RegisterAccountServiceServer(grpcServer, NewAccountService())
	asset_pb.RegisterAssetServiceServer(grpcServer, NewAssetService())
	battle_result_pb.RegisterBattleResultServiceServer(grpcServer, NewBattleResultService())
	battle_user_pb.RegisterBattleUserServiceServer(grpcServer, NewBattleUserService())
	character_pb.RegisterCharacterServiceServer(grpcServer, NewCharacterService())
	character_broadcast_pb.RegisterCharacterBroadcastServiceServer(grpcServer, NewCharacterBroadcastService())
	character_collection_pb.RegisterCharacterCollectionServiceServer(grpcServer, NewCharacterCollectionService())
	costume_pb.RegisterCostumeServiceServer(grpcServer, NewCostumeService())
	deck_pb.RegisterDeckServiceServer(grpcServer, NewDeckService())
	fate_card_pb.RegisterFateCardServiceServer(grpcServer, NewFateCardService())
	gacha_log_pb.RegisterGachaLogServiceServer(grpcServer, NewGachaLogService())
	item_pb.RegisterItemServiceServer(grpcServer, NewItemService())
	mail_pb.RegisterMailServiceServer(grpcServer, NewMailService())
	recipe_pb.RegisterRecipeServiceServer(grpcServer, NewRecipeService())
	shop_pb.RegisterShopServiceServer(grpcServer, NewShopService())
	story_pb.RegisterStoryServiceServer(grpcServer, NewStoryService())
	user_pb.RegisterUserServiceServer(grpcServer, NewUserService())

	if config.Config().Env != "production" {
		cheat_pb.RegisterCheatServiceServer(grpcServer, NewCheatService())
	}
}
