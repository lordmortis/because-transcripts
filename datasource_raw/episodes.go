// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datasource_raw

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Episode is an object representing the database table.
type Episode struct {
	ID          []byte      `boil:"id" json:"id" toml:"id" yaml:"id"`
	PodcastID   []byte      `boil:"podcast_id" json:"podcast_id" toml:"podcast_id" yaml:"podcast_id"`
	Number      null.Int64  `boil:"number" json:"number,omitempty" toml:"number" yaml:"number,omitempty"`
	Name        null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	AiredAt     time.Time   `boil:"aired_at" json:"aired_at" toml:"aired_at" yaml:"aired_at"`
	PatreonOnly null.Int64  `boil:"patreon_only" json:"patreon_only,omitempty" toml:"patreon_only" yaml:"patreon_only,omitempty"`

	R *episodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L episodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EpisodeColumns = struct {
	ID          string
	PodcastID   string
	Number      string
	Name        string
	AiredAt     string
	PatreonOnly string
}{
	ID:          "id",
	PodcastID:   "podcast_id",
	Number:      "number",
	Name:        "name",
	AiredAt:     "aired_at",
	PatreonOnly: "patreon_only",
}

var EpisodeTableColumns = struct {
	ID          string
	PodcastID   string
	Number      string
	Name        string
	AiredAt     string
	PatreonOnly string
}{
	ID:          "episodes.id",
	PodcastID:   "episodes.podcast_id",
	Number:      "episodes.number",
	Name:        "episodes.name",
	AiredAt:     "episodes.aired_at",
	PatreonOnly: "episodes.patreon_only",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var EpisodeWhere = struct {
	ID          whereHelper__byte
	PodcastID   whereHelper__byte
	Number      whereHelpernull_Int64
	Name        whereHelpernull_String
	AiredAt     whereHelpertime_Time
	PatreonOnly whereHelpernull_Int64
}{
	ID:          whereHelper__byte{field: "\"episodes\".\"id\""},
	PodcastID:   whereHelper__byte{field: "\"episodes\".\"podcast_id\""},
	Number:      whereHelpernull_Int64{field: "\"episodes\".\"number\""},
	Name:        whereHelpernull_String{field: "\"episodes\".\"name\""},
	AiredAt:     whereHelpertime_Time{field: "\"episodes\".\"aired_at\""},
	PatreonOnly: whereHelpernull_Int64{field: "\"episodes\".\"patreon_only\""},
}

// EpisodeRels is where relationship names are stored.
var EpisodeRels = struct {
	Podcast string
	Turns   string
}{
	Podcast: "Podcast",
	Turns:   "Turns",
}

// episodeR is where relationships are stored.
type episodeR struct {
	Podcast *Podcast  `boil:"Podcast" json:"Podcast" toml:"Podcast" yaml:"Podcast"`
	Turns   TurnSlice `boil:"Turns" json:"Turns" toml:"Turns" yaml:"Turns"`
}

// NewStruct creates a new relationship struct
func (*episodeR) NewStruct() *episodeR {
	return &episodeR{}
}

func (r *episodeR) GetPodcast() *Podcast {
	if r == nil {
		return nil
	}
	return r.Podcast
}

func (r *episodeR) GetTurns() TurnSlice {
	if r == nil {
		return nil
	}
	return r.Turns
}

// episodeL is where Load methods for each relationship are stored.
type episodeL struct{}

var (
	episodeAllColumns            = []string{"id", "podcast_id", "number", "name", "aired_at", "patreon_only"}
	episodeColumnsWithoutDefault = []string{"id", "podcast_id", "aired_at"}
	episodeColumnsWithDefault    = []string{"number", "name", "patreon_only"}
	episodePrimaryKeyColumns     = []string{"id"}
	episodeGeneratedColumns      = []string{}
)

type (
	// EpisodeSlice is an alias for a slice of pointers to Episode.
	// This should almost always be used instead of []Episode.
	EpisodeSlice []*Episode
	// EpisodeHook is the signature for custom Episode hook methods
	EpisodeHook func(context.Context, boil.ContextExecutor, *Episode) error

	episodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	episodeType                 = reflect.TypeOf(&Episode{})
	episodeMapping              = queries.MakeStructMapping(episodeType)
	episodePrimaryKeyMapping, _ = queries.BindMapping(episodeType, episodeMapping, episodePrimaryKeyColumns)
	episodeInsertCacheMut       sync.RWMutex
	episodeInsertCache          = make(map[string]insertCache)
	episodeUpdateCacheMut       sync.RWMutex
	episodeUpdateCache          = make(map[string]updateCache)
	episodeUpsertCacheMut       sync.RWMutex
	episodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var episodeAfterSelectHooks []EpisodeHook

var episodeBeforeInsertHooks []EpisodeHook
var episodeAfterInsertHooks []EpisodeHook

var episodeBeforeUpdateHooks []EpisodeHook
var episodeAfterUpdateHooks []EpisodeHook

var episodeBeforeDeleteHooks []EpisodeHook
var episodeAfterDeleteHooks []EpisodeHook

var episodeBeforeUpsertHooks []EpisodeHook
var episodeAfterUpsertHooks []EpisodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Episode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Episode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Episode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Episode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Episode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Episode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Episode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Episode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Episode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range episodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEpisodeHook registers your hook function for all future operations.
func AddEpisodeHook(hookPoint boil.HookPoint, episodeHook EpisodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		episodeAfterSelectHooks = append(episodeAfterSelectHooks, episodeHook)
	case boil.BeforeInsertHook:
		episodeBeforeInsertHooks = append(episodeBeforeInsertHooks, episodeHook)
	case boil.AfterInsertHook:
		episodeAfterInsertHooks = append(episodeAfterInsertHooks, episodeHook)
	case boil.BeforeUpdateHook:
		episodeBeforeUpdateHooks = append(episodeBeforeUpdateHooks, episodeHook)
	case boil.AfterUpdateHook:
		episodeAfterUpdateHooks = append(episodeAfterUpdateHooks, episodeHook)
	case boil.BeforeDeleteHook:
		episodeBeforeDeleteHooks = append(episodeBeforeDeleteHooks, episodeHook)
	case boil.AfterDeleteHook:
		episodeAfterDeleteHooks = append(episodeAfterDeleteHooks, episodeHook)
	case boil.BeforeUpsertHook:
		episodeBeforeUpsertHooks = append(episodeBeforeUpsertHooks, episodeHook)
	case boil.AfterUpsertHook:
		episodeAfterUpsertHooks = append(episodeAfterUpsertHooks, episodeHook)
	}
}

// One returns a single episode record from the query.
func (q episodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Episode, error) {
	o := &Episode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datasource_raw: failed to execute a one query for episodes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Episode records from the query.
func (q episodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (EpisodeSlice, error) {
	var o []*Episode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "datasource_raw: failed to assign all query results to Episode slice")
	}

	if len(episodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Episode records in the query.
func (q episodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to count episodes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q episodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "datasource_raw: failed to check if episodes exists")
	}

	return count > 0, nil
}

// Podcast pointed to by the foreign key.
func (o *Episode) Podcast(mods ...qm.QueryMod) podcastQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PodcastID),
	}

	queryMods = append(queryMods, mods...)

	return Podcasts(queryMods...)
}

// Turns retrieves all the turn's Turns with an executor.
func (o *Episode) Turns(mods ...qm.QueryMod) turnQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"turns\".\"episode_id\"=?", o.ID),
	)

	return Turns(queryMods...)
}

// LoadPodcast allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (episodeL) LoadPodcast(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEpisode interface{}, mods queries.Applicator) error {
	var slice []*Episode
	var object *Episode

	if singular {
		var ok bool
		object, ok = maybeEpisode.(*Episode)
		if !ok {
			object = new(Episode)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEpisode)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEpisode))
			}
		}
	} else {
		s, ok := maybeEpisode.(*[]*Episode)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEpisode)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEpisode))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &episodeR{}
		}
		if !queries.IsNil(object.PodcastID) {
			args = append(args, object.PodcastID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &episodeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.PodcastID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.PodcastID) {
				args = append(args, obj.PodcastID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`podcasts`),
		qm.WhereIn(`podcasts.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Podcast")
	}

	var resultSlice []*Podcast
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Podcast")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for podcasts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for podcasts")
	}

	if len(episodeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Podcast = foreign
		if foreign.R == nil {
			foreign.R = &podcastR{}
		}
		foreign.R.Episodes = append(foreign.R.Episodes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.PodcastID, foreign.ID) {
				local.R.Podcast = foreign
				if foreign.R == nil {
					foreign.R = &podcastR{}
				}
				foreign.R.Episodes = append(foreign.R.Episodes, local)
				break
			}
		}
	}

	return nil
}

// LoadTurns allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (episodeL) LoadTurns(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEpisode interface{}, mods queries.Applicator) error {
	var slice []*Episode
	var object *Episode

	if singular {
		var ok bool
		object, ok = maybeEpisode.(*Episode)
		if !ok {
			object = new(Episode)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEpisode)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEpisode))
			}
		}
	} else {
		s, ok := maybeEpisode.(*[]*Episode)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEpisode)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEpisode))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &episodeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &episodeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`turns`),
		qm.WhereIn(`turns.episode_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load turns")
	}

	var resultSlice []*Turn
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice turns")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on turns")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for turns")
	}

	if len(turnAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Turns = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &turnR{}
			}
			foreign.R.Episode = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.EpisodeID) {
				local.R.Turns = append(local.R.Turns, foreign)
				if foreign.R == nil {
					foreign.R = &turnR{}
				}
				foreign.R.Episode = local
				break
			}
		}
	}

	return nil
}

// SetPodcast of the episode to the related item.
// Sets o.R.Podcast to related.
// Adds o to related.R.Episodes.
func (o *Episode) SetPodcast(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Podcast) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"episodes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"podcast_id"}),
		strmangle.WhereClause("\"", "\"", 0, episodePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.PodcastID, related.ID)
	if o.R == nil {
		o.R = &episodeR{
			Podcast: related,
		}
	} else {
		o.R.Podcast = related
	}

	if related.R == nil {
		related.R = &podcastR{
			Episodes: EpisodeSlice{o},
		}
	} else {
		related.R.Episodes = append(related.R.Episodes, o)
	}

	return nil
}

// AddTurns adds the given related objects to the existing relationships
// of the episode, optionally inserting them as new records.
// Appends related to o.R.Turns.
// Sets related.R.Episode appropriately.
func (o *Episode) AddTurns(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Turn) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.EpisodeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"turns\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"episode_id"}),
				strmangle.WhereClause("\"", "\"", 0, turnPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.EpisodeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &episodeR{
			Turns: related,
		}
	} else {
		o.R.Turns = append(o.R.Turns, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &turnR{
				Episode: o,
			}
		} else {
			rel.R.Episode = o
		}
	}
	return nil
}

// Episodes retrieves all the records using an executor.
func Episodes(mods ...qm.QueryMod) episodeQuery {
	mods = append(mods, qm.From("\"episodes\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"episodes\".*"})
	}

	return episodeQuery{q}
}

// FindEpisode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEpisode(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*Episode, error) {
	episodeObj := &Episode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"episodes\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, episodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datasource_raw: unable to select from episodes")
	}

	if err = episodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return episodeObj, err
	}

	return episodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Episode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("datasource_raw: no episodes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(episodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	episodeInsertCacheMut.RLock()
	cache, cached := episodeInsertCache[key]
	episodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			episodeAllColumns,
			episodeColumnsWithDefault,
			episodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(episodeType, episodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(episodeType, episodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"episodes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"episodes\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "datasource_raw: unable to insert into episodes")
	}

	if !cached {
		episodeInsertCacheMut.Lock()
		episodeInsertCache[key] = cache
		episodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Episode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Episode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	episodeUpdateCacheMut.RLock()
	cache, cached := episodeUpdateCache[key]
	episodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			episodeAllColumns,
			episodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("datasource_raw: unable to update episodes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"episodes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, episodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(episodeType, episodeMapping, append(wl, episodePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to update episodes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by update for episodes")
	}

	if !cached {
		episodeUpdateCacheMut.Lock()
		episodeUpdateCache[key] = cache
		episodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q episodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to update all for episodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to retrieve rows affected for episodes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EpisodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("datasource_raw: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), episodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"episodes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, episodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to update all in episode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to retrieve rows affected all in update all episode")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Episode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("datasource_raw: no episodes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(episodeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	episodeUpsertCacheMut.RLock()
	cache, cached := episodeUpsertCache[key]
	episodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			episodeAllColumns,
			episodeColumnsWithDefault,
			episodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			episodeAllColumns,
			episodePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("datasource_raw: unable to upsert episodes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(episodePrimaryKeyColumns))
			copy(conflict, episodePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"episodes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(episodeType, episodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(episodeType, episodeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "datasource_raw: unable to upsert episodes")
	}

	if !cached {
		episodeUpsertCacheMut.Lock()
		episodeUpsertCache[key] = cache
		episodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Episode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Episode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("datasource_raw: no Episode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), episodePrimaryKeyMapping)
	sql := "DELETE FROM \"episodes\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to delete from episodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by delete for episodes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q episodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("datasource_raw: no episodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to delete all from episodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by deleteall for episodes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EpisodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(episodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), episodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"episodes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, episodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to delete all from episode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by deleteall for episodes")
	}

	if len(episodeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Episode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEpisode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EpisodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EpisodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), episodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"episodes\".* FROM \"episodes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, episodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "datasource_raw: unable to reload all in EpisodeSlice")
	}

	*o = slice

	return nil
}

// EpisodeExists checks if the Episode row exists.
func EpisodeExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"episodes\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "datasource_raw: unable to check if episodes exists")
	}

	return exists, nil
}
