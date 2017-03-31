package cli

import (
	"../client"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// KeepEventEventsCommand is the command line data structure for the keep event action of events
	KeepEventEventsCommand struct {
		// イベントID
		EventID int
		// キープ操作
		IsKeep      string
		PrettyPrint bool
	}

	// ListEventsCommand is the command line data structure for the list action of events
	ListEventsCommand struct {
		ID string
		// ページ(1->2->3->4)
		Page int
		// キーワード検索
		Q string
		// ソート
		Sort        string
		PrettyPrint bool
	}

	// CreateGenresCommand is the command line data structure for the create action of genres
	CreateGenresCommand struct {
		// ジャンル名
		Name        string
		PrettyPrint bool
	}

	// FollowGenreGenresCommand is the command line data structure for the follow genre action of genres
	FollowGenreGenresCommand struct {
		// ジャンルID
		GenreID     int
		PrettyPrint bool
	}

	// ListGenresCommand is the command line data structure for the list action of genres
	ListGenresCommand struct {
		// ジャンル名検索に使うキーワード
		Q           string
		PrettyPrint bool
	}

	// PrefFollowPrefsCommand is the command line data structure for the pref follow action of prefs
	PrefFollowPrefsCommand struct {
		// 都道府県ID
		PrefID      int
		PrettyPrint bool
	}

	// AccountCreateUsersCommand is the command line data structure for the account create action of users
	AccountCreateUsersCommand struct {
		// アプリのバージョン
		ClientVersion string
		// メールアドレス
		Email string
		// 識別子(android:Android_ID, ios:IDFV)
		Identifier string
		// OSとバージョン
		Platform    string
		PrettyPrint bool
	}

	// TmpAccountCreateUsersCommand is the command line data structure for the tmp account create action of users
	TmpAccountCreateUsersCommand struct {
		// アプリのバージョン
		ClientVersion string
		// 識別子(android:Android_ID, ios:IDFV)
		Identifier string
		// OSとバージョン
		Platform    string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "accountCreate",
		Short: `正規ユーザーの作成`,
	}
	tmp1 := new(AccountCreateUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/v2/users/new"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "create",
		Short: `ジャンルの新規作成`,
	}
	tmp2 := new(CreateGenresCommand)
	sub = &cobra.Command{
		Use:   `genres ["/api/v2/genres/new"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "followGenre",
		Short: `ジャンルお気に入り操作`,
	}
	tmp3 := new(FollowGenreGenresCommand)
	sub = &cobra.Command{
		Use:   `genres [("/api/v2/genres/GENREID/follow"|"/api/v2/genres/GENREID/follow")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "keepEvent",
		Short: `イベントのお気に入り操作`,
	}
	tmp4 := new(KeepEventEventsCommand)
	sub = &cobra.Command{
		Use:   `events ["/api/v2/events/EVENTID/keep"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp5 := new(ListEventsCommand)
	sub = &cobra.Command{
		Use:   `events [("/api/v2/events/genre/ID"|"/api/v2/events/new"|"/api/v2/events/keep"|"/api/v2/events/nokeep"|"/api/v2/events/popular"|"/api/v2/events/recommend")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(ListGenresCommand)
	sub = &cobra.Command{
		Use:   `genres ["/api/v2/genres"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "prefFollow",
		Short: `ジャンルお気に入り操作`,
	}
	tmp7 := new(PrefFollowPrefsCommand)
	sub = &cobra.Command{
		Use:   `prefs [("/api/v2/prefs/PREFID/follow"|"/api/v2/prefs/PREFID/follow")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "tmpAccountCreate",
		Short: `一時ユーザーの作成`,
	}
	tmp8 := new(TmpAccountCreateUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/v2/users/tmp"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the KeepEventEventsCommand command.
func (cmd *KeepEventEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/v2/events/%v/keep", cmd.EventID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp9 *bool
	if cmd.IsKeep != "" {
		var err error
		tmp9, err = boolVal(cmd.IsKeep)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--isKeep", "err", err)
			return err
		}
	}
	resp, err := c.KeepEventEvents(ctx, path, *tmp9)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *KeepEventEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, `イベントID`)
	var isKeep string
	cc.Flags().StringVar(&cmd.IsKeep, "isKeep", isKeep, `キープ操作`)
}

// Run makes the HTTP request corresponding to the ListEventsCommand command.
func (cmd *ListEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/events/new"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListEvents(ctx, path, intFlagVal("page", cmd.Page), stringFlagVal("q", cmd.Q), stringFlagVal("sort", cmd.Sort))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
	var page int
	cc.Flags().IntVar(&cmd.Page, "page", page, `ページ(1->2->3->4)`)
	var q string
	cc.Flags().StringVar(&cmd.Q, "q", q, `キーワード検索`)
	var sort string
	cc.Flags().StringVar(&cmd.Sort, "sort", sort, `ソート`)
}

// Run makes the HTTP request corresponding to the CreateGenresCommand command.
func (cmd *CreateGenresCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/genres/new"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateGenres(ctx, path, cmd.Name)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateGenresCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var name string
	cc.Flags().StringVar(&cmd.Name, "name", name, `ジャンル名`)
}

// Run makes the HTTP request corresponding to the FollowGenreGenresCommand command.
func (cmd *FollowGenreGenresCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/v2/genres/%v/follow", cmd.GenreID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.FollowGenreGenres(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *FollowGenreGenresCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var genreID int
	cc.Flags().IntVar(&cmd.GenreID, "genreID", genreID, `ジャンルID`)
}

// Run makes the HTTP request corresponding to the ListGenresCommand command.
func (cmd *ListGenresCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/genres"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListGenres(ctx, path, stringFlagVal("q", cmd.Q))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListGenresCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var q string
	cc.Flags().StringVar(&cmd.Q, "q", q, `ジャンル名検索に使うキーワード`)
}

// Run makes the HTTP request corresponding to the PrefFollowPrefsCommand command.
func (cmd *PrefFollowPrefsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/v2/prefs/%v/follow", cmd.PrefID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.PrefFollowPrefs(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *PrefFollowPrefsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var prefID int
	cc.Flags().IntVar(&cmd.PrefID, "prefID", prefID, `都道府県ID`)
}

// Run makes the HTTP request corresponding to the AccountCreateUsersCommand command.
func (cmd *AccountCreateUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/users/new"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AccountCreateUsers(ctx, path, cmd.ClientVersion, cmd.Email, cmd.Identifier, cmd.Platform)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AccountCreateUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var clientVersion string
	cc.Flags().StringVar(&cmd.ClientVersion, "client_version", clientVersion, `アプリのバージョン`)
	var email string
	cc.Flags().StringVar(&cmd.Email, "email", email, `メールアドレス`)
	var identifier string
	cc.Flags().StringVar(&cmd.Identifier, "identifier", identifier, `識別子(android:Android_ID, ios:IDFV)`)
	var platform string
	cc.Flags().StringVar(&cmd.Platform, "platform", platform, `OSとバージョン`)
}

// Run makes the HTTP request corresponding to the TmpAccountCreateUsersCommand command.
func (cmd *TmpAccountCreateUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/users/tmp"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.TmpAccountCreateUsers(ctx, path, cmd.ClientVersion, cmd.Identifier, cmd.Platform)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TmpAccountCreateUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var clientVersion string
	cc.Flags().StringVar(&cmd.ClientVersion, "client_version", clientVersion, `アプリのバージョン`)
	var identifier string
	cc.Flags().StringVar(&cmd.Identifier, "identifier", identifier, `識別子(android:Android_ID, ios:IDFV)`)
	var platform string
	cc.Flags().StringVar(&cmd.Platform, "platform", platform, `OSとバージョン`)
}
