// GoNeSh - Go Neural Shell
// AI時代に特化したエンジニア向け次世代シェル環境
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ousiass/GoNeSh/internal/core"
	"github.com/ousiass/GoNeSh/internal/errors"
	"github.com/ousiass/GoNeSh/pkg/config"
)

const version = "0.1.0"

func main() {
	// 設定を読み込む
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("%v\n", errors.Wrap(errors.E1001, err))
		fmt.Println("ヒント: gonesh --init で初期設定を作成できます")
		os.Exit(1)
	}

	// アプリケーションを初期化
	app := core.NewApp(cfg)

	// Bubbleteaプログラムを開始
	p := tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("%v\n", errors.Wrap(errors.E2001, err))
		os.Exit(1)
	}
}
