package v1

import (
	bankgold "napoleon-email/src/app/application/bank_gold"
	groupnapoleon "napoleon-email/src/app/application/group_napoleon"
	"napoleon-email/src/app/application/mine"
	"napoleon-email/src/app/application/napoleon"
	bankgoldhandler "napoleon-email/src/app/http/handler/bank_gold_handler"
	groupnapoleonhandler "napoleon-email/src/app/http/handler/group_napoleon_handler"
	napoleonhandler "napoleon-email/src/app/http/handler/napoleon_handler"
	napoleonminehandler "napoleon-email/src/app/http/handler/napoleon_mine_handler"
	"napoleon-email/src/app/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func RouterApiV1(app *fiber.App, c *infrastructure.Kernel) {
	napoUseCase := napoleon.NewNapoEmailApplicationImp(c.GetContactNapoleonRepository())
	groupUseCase := groupnapoleon.NewNapoGroupEmailApplicationImp(c.GetContactGroupNapoleonRepository())
	napoMinerUseCase := mine.NewNapoMineEmailApplicationImp(c.GetContactNapoleonMineRepository())
	bankGoldUseCase := bankgold.NewBankGoldEmailApplicationImp(c.GetContactBankGoldRepository())

	napoHandler := napoleonhandler.NewNapoleonHandler(napoUseCase)
	groupNapoHandler := groupnapoleonhandler.NewGroupNapoleonHandler(groupUseCase)
	napoMineHandler := napoleonminehandler.NewNapoleonMineHandler(napoMinerUseCase)
	bankGoldHandler := bankgoldhandler.NewBankGoldHandler(bankGoldUseCase)

	http := app.Group("/api/v1")
	napo := http.Group("/napoleon")
	groupNapo := http.Group("/group-napoleon")
	bankGold := http.Group("/bank-gold")
	napoMiner := http.Group("/napoleon-mine")


	napo.Post("/contact", napoHandler.CreateNapoEmail)
	groupNapo.Post("/contact", groupNapoHandler.CreateGroupNapoEmail)
	bankGold.Post("/contact", bankGoldHandler.CreateBankGoldEmail)
	napoMiner.Post("/contact", napoMineHandler.CreateNapoMineEmail)
}
