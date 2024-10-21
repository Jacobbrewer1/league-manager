package cli

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/jacobbrewer1/league-manager/pkg/codegen/apis/api"
)

func (s *service) Run() error {
	//form := huh.NewForm(
	//	huh.NewGroup(
	//		// Ask the user for a base burger and toppings.
	//		huh.NewSelect[string]().
	//			Title("Choose your burger").
	//			Options(
	//				huh.NewOption("Charmburger Classic", "classic"),
	//				huh.NewOption("Chickwich", "chickwich"),
	//				huh.NewOption("Fishburger", "fishburger"),
	//				huh.NewOption("Charmpossible™ Burger", "charmpossible"),
	//			).
	//			Value(&burger), // store the chosen option in the "burger" variable
	//
	//		// Let the user select multiple toppings.
	//		huh.NewMultiSelect[string]().
	//			Title("Toppings").
	//			Options(
	//				huh.NewOption("Lettuce", "lettuce").Selected(true),
	//				huh.NewOption("Tomatoes", "tomatoes").Selected(true),
	//				huh.NewOption("Jalapeños", "jalapeños"),
	//				huh.NewOption("Cheese", "cheese"),
	//				huh.NewOption("Vegan Cheese", "vegan cheese"),
	//				huh.NewOption("Nutella", "nutella"),
	//			).
	//			Limit(4). // there’s a 4 topping limit!
	//			Value(&toppings),
	//
	//		// Option values in selects and multi selects can be any type you
	//		// want. We’ve been recording strings above, but here we’ll store
	//		// answers as integers. Note the generic "[int]" directive below.
	//		huh.NewSelect[int]().
	//			Title("How much Charm Sauce do you want?").
	//			Options(
	//				huh.NewOption("None", 0),
	//				huh.NewOption("A little", 1),
	//				huh.NewOption("A lot", 2),
	//			).
	//			Value(&sauceLevel),
	//	),
	//
	//	// Gather some final details about the order.
	//	huh.NewGroup(
	//		huh.NewInput().
	//			Title("What’s your name?").
	//			Value(&name).
	//			// Validating fields is easy. The form will mark erroneous fields
	//			// and display error messages accordingly.
	//			Validate(func(str string) error {
	//				if str == "Frank" {
	//					return errors.New("sorry, we don’t serve customers named Frank")
	//				}
	//				return nil
	//			}),
	//
	//		huh.NewText().
	//			Title("Special Instructions").
	//			CharLimit(400).
	//			Value(&instructions),
	//
	//		huh.NewConfirm().
	//			Title("Would you like 15% off?").
	//			Value(&discount),
	//	),
	//)

	clientErrs := make(chan error, 1)
	playersChan := make(chan *api.PlayersResponse, 1)
	err := spinner.New().
		Title("Making your burger...").
		Action(s.getPlayers(playersChan, clientErrs)).
		Run()
	if err != nil {
		return fmt.Errorf("error getting players: %w", err)
	} else if len(clientErrs) > 0 {
		return fmt.Errorf("error getting players: %s", <-clientErrs)
	}
	players := <-playersChan
	close(playersChan)

	fmt.Println("Players:")
	for _, player := range players.Players {
		fmt.Println(*player.FirstName)
	}

	fmt.Println("Order up!")

	form := huh.NewForm()

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *service) getPlayers(dest chan *api.PlayersResponse, destErrs chan error) func() {
	return func() {
		params := new(api.GetPlayersParams)
		resp, err := s.r.GetPlayers(params)
		if err != nil {
			destErrs <- fmt.Errorf("error getting players: %w", err)
			return
		}
		close(destErrs)

		dest <- resp
	}
}
