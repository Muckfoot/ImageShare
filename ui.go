package main

import (
	"strconv"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

var (
	appGeoSettings AppGeoSettings
)
var (
	ui_tools_list           *widgets.QListWidget
	ui_image_preview        *widgets.QLabel
	ui_session_history_list *widgets.QTableWidget
	pix                     *gui.QPixmap
)

func NewImageShareForm() *widgets.QWidget {

	widget := widgets.NewQWidget(nil, 0)

	loader := uitools.NewQUiLoader(nil)
	file := core.NewQFile2(root + "ImageShareUi2.ui")

	file.Open(core.QIODevice__ReadOnly)
	formWidget := loader.Load(file, widget)
	file.Close()

	loadJSON(root+"appGeo.json", &appGeoSettings)

	widget.SetGeometry2(
		appGeoSettings.MainWindowGeo.Pos[0],
		appGeoSettings.MainWindowGeo.Pos[1],
		appGeoSettings.MainWindowGeo.Size[0],
		appGeoSettings.MainWindowGeo.Size[1],
	)

	ui_tools_list = widgets.NewQListWidgetFromPointer(
		widget.FindChild(
			"ToolsListView", core.Qt__FindChildrenRecursively).Pointer())
	ui_image_preview = widgets.NewQLabelFromPointer(
		widget.FindChild(
			"ImagePreview", core.Qt__FindChildrenRecursively).Pointer())
	ui_session_history_list = widgets.NewQTableWidgetFromPointer(
		widget.FindChild(
			"SessionHistoryList", core.Qt__FindChildrenRecursively).Pointer())

	ui_splitter := widgets.NewQSplitterFromPointer(
		widget.FindChild(
			"splitter", core.Qt__FindChildrenRecursively).Pointer())
	ui_frame := widgets.NewQFrameFromPointer(
		widget.FindChild(
			"ImagePreviewFrame", core.Qt__FindChildrenRecursively).Pointer())

	ui_tools_list.Resize2(
		appGeoSettings.ToolsListView.Width,
		appGeoSettings.ToolsListView.Height)

	ui_frame.Resize2(
		appGeoSettings.ImagePreviewFrame.Width,
		appGeoSettings.ImagePreviewFrame.Height)

	ui_session_history_list.Resize2(
		appGeoSettings.SessionHistoryList.Width,
		appGeoSettings.SessionHistoryList.Height)

	pix = gui.NewQPixmap5(root+"ss.png", "", 0)
	ui_image_preview.SetScaledContents(false)
	ui_image_preview.SetPixmap(
		pix.Scaled(
			ui_frame.Size(),
			core.Qt__KeepAspectRatio,
			core.Qt__SmoothTransformation))

	ui_session_history_list.ResizeColumnsToContents()
	ui_session_history_list.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(formWidget, 0, 0)
	widget.SetLayout(layout)

	widget.SetWindowTitle("ImageShare")

	widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		newFrameSize := ui_frame.Size()
		ui_image_preview.SetPixmap(
			pix.Scaled(
				newFrameSize,
				core.Qt__KeepAspectRatio,
				core.Qt__SmoothTransformation))
	})

	ui_splitter.ConnectSplitterMoved(func(pos int, index int) {

		updateScaling(ui_frame.Size())
	})

	ui_session_history_list.ConnectCellClicked(func(row int, column int) {
		item := ui_session_history_list.Item(row, 0)
		pix.Load(path+item.Text(), "", 0)
		ui_image_preview.SetPixmap(pix.FromImage(pix.ToImage(), 0))
		updateScaling(ui_frame.Size())
	})

	widget.ConnectCloseEvent(func(event *gui.QCloseEvent) {

		geo := widget.Geometry()

		appGeoSettings.MainWindowGeo.Pos = []int{geo.X(), geo.Y()}
		appGeoSettings.MainWindowGeo.Size = []int{geo.Width(), geo.Height()}

		appGeoSettings.ToolsListView.Height = ui_tools_list.Height()
		appGeoSettings.ToolsListView.Width = ui_tools_list.Width()

		appGeoSettings.SessionHistoryList.Height = ui_session_history_list.Height()
		appGeoSettings.SessionHistoryList.Width = ui_session_history_list.Width()

		appGeoSettings.ImagePreviewFrame.Height = ui_frame.Height()
		appGeoSettings.ImagePreviewFrame.Width = ui_frame.Width()

		saveJSON(appGeoSettings, root+"appGeo.json")

	})

	return widget
}

func updateHistoryList(link, fname string, size int) {
	params := []string{fname, strconv.Itoa(size), link}

	ui_session_history_list.SetRowCount(ui_session_history_list.RowCount() + 1)
	for i := 0; i < 3; i++ {
		ui_session_history_list.SetItem(
			ui_session_history_list.RowCount()-1, i, widgets.NewQTableWidgetItem2(params[i], 0))
	}
}

func updateScaling(newFrameSize *core.QSize) {
	time.Sleep(time.Millisecond * 20)
	ui_image_preview.SetPixmap(pix.Scaled(
		newFrameSize,
		core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation))
}
