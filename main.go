package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

var soldierDropdown *widget.Select
var peasantDropdown *widget.Select
var pawnDropdown *widget.Select
var rookDropdown *widget.Select
var knightDropdown *widget.Select
var bishopDropdown *widget.Select
var catapultDropdown *widget.Select
var courtesanDropdown *widget.Select
var chamberlainDropdown *widget.Select
var heraldDropdown *widget.Select
var inquisitorDropdown *widget.Select
var lancerDropdown *widget.Select
var pontiffDropdown *widget.Select
var thiefDropdown *widget.Select
var towerDropdown *widget.Select
var queenDropdown *widget.Select
var jesterDropdown *widget.Select
var kingDropdown *widget.Select
var regentDropdown *widget.Select
var rankIDropdowns []*widget.Select
var rankIIDropdowns []*widget.Select
var rankIIIDropdowns []*widget.Select

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Faerie Chess Counter")

	pawnCombo := []string{"4", "5", "6", "7", "8"}
	nilToTwoCombo := []string{"0", "1", "2"}
	nilToOneCombo := []string{"0", "1"}
	difficultyCombo := []string{"Beginner", "Intermediate", "Expert"}

	//a bad solution to the comboBox sizing using this layout manager
	spacer := widget.NewLabel("")
	resultLabel := widget.NewLabel("Total Points: ")
	pointsRemainingLabel := widget.NewLabel("Points Remaining: ")

	rankIRemainingLabel := widget.NewLabel("Rank I Remaining: ")
	rankIIRemainingLabel := widget.NewLabel("Rank II Remaining: ")
	rankIIIRemainingLabel := widget.NewLabel("Rank III Remaining: ")

	//rankI labels and dropdowns
	pawnLabel := widget.NewLabel("Pawns:")
	pawnDropdown = widget.NewSelect(pawnCombo, func(selected string) {
		updateRankRemainingLabel(rankIDropdowns, rankIRemainingLabel, 8)
	})
	pawnDropdown.SetSelectedIndex(0)

	peasantLabel := widget.NewLabel("Peasants:")
	peasantDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIDropdowns, rankIRemainingLabel, 8)
	})
	peasantDropdown.SetSelectedIndex(0)

	soldierLabel := widget.NewLabel("Soldiers:")
	soldierDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIDropdowns, rankIRemainingLabel, 8)
	})
	soldierDropdown.SetSelectedIndex(0)

	//rankII labels and dropdowns
	rookLabel := widget.NewLabel("Rooks:")
	rookDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	rookDropdown.SetSelectedIndex(0)

	knightLabel := widget.NewLabel("Knights:")
	knightDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	knightDropdown.SetSelectedIndex(0)

	bishopLabel := widget.NewLabel("Bishops:")
	bishopDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	bishopDropdown.SetSelectedIndex(0)

	catapultLabel := widget.NewLabel("Catapult:")
	catapultDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	catapultDropdown.SetSelectedIndex(0)

	courtesanLabel := widget.NewLabel("Courtesan:")
	courtesanDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	courtesanDropdown.SetSelectedIndex(0)

	chamberlainLabel := widget.NewLabel("Chamberlain:")
	chamberlainDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	chamberlainDropdown.SetSelectedIndex(0)

	heraldLabel := widget.NewLabel("Herald:")
	heraldDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	heraldDropdown.SetSelectedIndex(0)

	inquisitorLabel := widget.NewLabel("Inquisitor:")
	inquisitorDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	inquisitorDropdown.SetSelectedIndex(0)

	lancerLabel := widget.NewLabel("Lancer:")
	lancerDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	lancerDropdown.SetSelectedIndex(0)

	pontiffLabel := widget.NewLabel("Pontiff:")
	pontiffDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	pontiffDropdown.SetSelectedIndex(0)

	thiefLabel := widget.NewLabel("Thief:")
	thiefDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	thiefDropdown.SetSelectedIndex(0)

	towerLabel := widget.NewLabel("Tower:")
	towerDropdown = widget.NewSelect(nilToTwoCombo, func(selected string) {
		updateRankRemainingLabel(rankIIDropdowns, rankIIRemainingLabel, 6)
	})
	towerDropdown.SetSelectedIndex(0)

	//rankIII labels and dropdowns
	queenLabel := widget.NewLabel("Queen:")
	queenDropdown = widget.NewSelect(nilToOneCombo, func(selected string) {
		updateRankRemainingLabel(rankIIIDropdowns, rankIIIRemainingLabel, 2)
	})
	queenDropdown.SetSelectedIndex(0)

	kingLabel := widget.NewLabel("King:")
	kingDropdown = widget.NewSelect(nilToOneCombo, func(selected string) {
		updateRankRemainingLabel(rankIIIDropdowns, rankIIIRemainingLabel, 2)
	})
	kingDropdown.SetSelectedIndex(0)

	jesterLabel := widget.NewLabel("Jester:")
	jesterDropdown = widget.NewSelect(nilToOneCombo, func(selected string) {
		updateRankRemainingLabel(rankIIIDropdowns, rankIIIRemainingLabel, 2)
	})
	jesterDropdown.SetSelectedIndex(0)

	regentLabel := widget.NewLabel("Regent:")
	regentDropdown = widget.NewSelect(nilToOneCombo, func(selected string) {
		updateRankRemainingLabel(rankIIIDropdowns, rankIIIRemainingLabel, 2)
	})
	regentDropdown.SetSelectedIndex(0)

	//rankI dropdown lists for functions
	rankIDropdowns = append(rankIDropdowns, pawnDropdown, peasantDropdown, soldierDropdown)

	//rankII dropdown lists for functions
	rankIIDropdowns = append(rankIIDropdowns, rookDropdown, knightDropdown, bishopDropdown, catapultDropdown, courtesanDropdown, chamberlainDropdown, heraldDropdown, inquisitorDropdown, lancerDropdown, pontiffDropdown, thiefDropdown, towerDropdown)

	//rankIII dropdown lists for functions
	rankIIIDropdowns = append(rankIIIDropdowns, kingDropdown, regentDropdown, queenDropdown, jesterDropdown)

	difficultyLabel := widget.NewLabel("Difficulty:")
	difficultyDropdown := widget.NewSelect(difficultyCombo, func(selected string) {

	})
	difficultyDropdown.SetSelectedIndex(0)

	//create calculate function here
	calculateButton := widget.NewButton("Calculate", func() {
		pawnsSelected, err := strconv.Atoi(pawnDropdown.Selected)
		peasantsSelected, err := strconv.Atoi(peasantDropdown.Selected)
		soldiersSelected, err := strconv.Atoi(soldierDropdown.Selected)
		rooksSelected, err := strconv.Atoi(rookDropdown.Selected)
		knightsSelected, err := strconv.Atoi(knightDropdown.Selected)
		bishopsSelected, err := strconv.Atoi(bishopDropdown.Selected)
		catapultSelected, err := strconv.Atoi(catapultDropdown.Selected)
		chamberlainSelected, err := strconv.Atoi(chamberlainDropdown.Selected)
		courtesanSelected, err := strconv.Atoi(courtesanDropdown.Selected)
		heraldSelected, err := strconv.Atoi(heraldDropdown.Selected)
		inquisitorSelected, err := strconv.Atoi(inquisitorDropdown.Selected)
		lancerSelected, err := strconv.Atoi(lancerDropdown.Selected)
		pontiffSelected, err := strconv.Atoi(pontiffDropdown.Selected)
		thiefSelected, err := strconv.Atoi(thiefDropdown.Selected)
		towerSelected, err := strconv.Atoi(towerDropdown.Selected)
		queenSelected, err := strconv.Atoi(queenDropdown.Selected)
		jesterSelected, err := strconv.Atoi(jesterDropdown.Selected)
		kingSelected, err := strconv.Atoi(kingDropdown.Selected)
		regentSelected, err := strconv.Atoi(regentDropdown.Selected)
		difficultySelected := difficultyDropdown.Selected

		if err != nil {
			panic(err)
		}

		var difficultyPoints int
		switch difficultySelected {
		case "Beginner":
			difficultyPoints = 65
		case "Intermediate":
			difficultyPoints = 70
		case "Advanced":
			difficultyPoints = 75
		default:
			difficultyPoints = 65
		}

		//use a dictionary here for point value
		totalPoints := (pawnsSelected * 1) + (peasantsSelected * 2) + (soldiersSelected * 3) + (rooksSelected * 9) + (knightsSelected * 4) + (bishopsSelected * 6) + (catapultSelected * 3) + (chamberlainSelected * 6) + (courtesanSelected * 6) + (heraldSelected * 6) + (inquisitorSelected * 8) + (lancerSelected * 5) + (pontiffSelected * 8) + (thiefSelected * 5) + (towerSelected * 10) + (queenSelected * 12) + (jesterSelected * 12) + (kingSelected * 0) + (regentSelected * 15)

		//total points  selected
		totalPointsString := strconv.Itoa(totalPoints)
		resultLabel.SetText("Total Points Selected: " + totalPointsString)
		//points remaining
		pointsRemaining := difficultyPoints - totalPoints
		pointsRemainingString := strconv.Itoa(pointsRemaining)
		pointsRemainingLabel.SetText("Points Remaining: " + pointsRemainingString)
	})

	rank1Container := container.New(layout.NewGridLayout(2),
		pawnLabel,
		pawnDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		peasantLabel,
		peasantDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		soldierLabel,
		soldierDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		rankIRemainingLabel,
		spacer,
		rankIIRemainingLabel,
		spacer,
		rankIIIRemainingLabel,
		spacer,
	)
	rankIRemainingLabel.SetText("Rank I Remaining: 4")

	rank2Container := container.New(layout.NewGridLayout(2),
		rookLabel,
		rookDropdown,
		knightLabel,
		knightDropdown,
		bishopLabel,
		bishopDropdown,
		catapultLabel,
		catapultDropdown,
		chamberlainLabel,
		chamberlainDropdown,
		courtesanLabel,
		courtesanDropdown,
		heraldLabel,
		heraldDropdown,
		inquisitorLabel,
		inquisitorDropdown,
		lancerLabel,
		lancerDropdown,
		pontiffLabel,
		pontiffDropdown,
		thiefLabel,
		thiefDropdown,
		towerLabel,
		towerDropdown,
		spacer,
		spacer,
		difficultyLabel,
		difficultyDropdown,
	)

	rank3Container := container.New(layout.NewGridLayout(2),
		kingLabel,
		kingDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		queenLabel,
		queenDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		regentLabel,
		regentDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		jesterLabel,
		jesterDropdown,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		spacer,
		resultLabel,
		calculateButton,
		pointsRemainingLabel,
	)

	mainContainer := container.New(layout.NewGridLayout(3),
		rank1Container,
		rank2Container,
		rank3Container,
	)

	myWindow.SetContent(mainContainer)
	myWindow.ShowAndRun()
}

func updateRankRemainingLabel(dropdowns []*widget.Select, label *widget.Label, maxPieces int) {
	var totalSelected int

	for _, dropdown := range dropdowns {
		if dropdown == nil {
			return
		}

		selected, err := strconv.Atoi(dropdown.Selected)

		if err != nil {
			panic(err)
		}

		totalSelected += selected
	}

	var labelText string
	switch maxPieces {
	case 8:
		labelText = "Rank I Remaining: "
	case 6:
		labelText = "Rank II Remaining: "
	case 2:
		labelText = "Rank III Remaining: "
	default:
		labelText = "Remaining: "
	}

	rankLeft := maxPieces - totalSelected
	rankLeftString := strconv.Itoa(rankLeft)
	label.SetText(labelText + rankLeftString)
}
