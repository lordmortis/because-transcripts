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

// Turn is an object representing the database table.
type Turn struct {
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	EpisodeID  string    `boil:"episode_id" json:"episode_id" toml:"episode_id" yaml:"episode_id"`
	SequenceNo int       `boil:"sequence_no" json:"sequence_no" toml:"sequence_no" yaml:"sequence_no"`
	StartTime  null.Int  `boil:"start_time" json:"start_time,omitempty" toml:"start_time" yaml:"start_time,omitempty"`
	EndTime    null.Int  `boil:"end_time" json:"end_time,omitempty" toml:"end_time" yaml:"end_time,omitempty"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *turnR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L turnL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TurnColumns = struct {
	ID         string
	EpisodeID  string
	SequenceNo string
	StartTime  string
	EndTime    string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	EpisodeID:  "episode_id",
	SequenceNo: "sequence_no",
	StartTime:  "start_time",
	EndTime:    "end_time",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var TurnTableColumns = struct {
	ID         string
	EpisodeID  string
	SequenceNo string
	StartTime  string
	EndTime    string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "turns.id",
	EpisodeID:  "turns.episode_id",
	SequenceNo: "turns.sequence_no",
	StartTime:  "turns.start_time",
	EndTime:    "turns.end_time",
	CreatedAt:  "turns.created_at",
	UpdatedAt:  "turns.updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var TurnWhere = struct {
	ID         whereHelperstring
	EpisodeID  whereHelperstring
	SequenceNo whereHelperint
	StartTime  whereHelpernull_Int
	EndTime    whereHelpernull_Int
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperstring{field: "\"turns\".\"id\""},
	EpisodeID:  whereHelperstring{field: "\"turns\".\"episode_id\""},
	SequenceNo: whereHelperint{field: "\"turns\".\"sequence_no\""},
	StartTime:  whereHelpernull_Int{field: "\"turns\".\"start_time\""},
	EndTime:    whereHelpernull_Int{field: "\"turns\".\"end_time\""},
	CreatedAt:  whereHelpertime_Time{field: "\"turns\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"turns\".\"updated_at\""},
}

// TurnRels is where relationship names are stored.
var TurnRels = struct {
	Episode    string
	Utterances string
}{
	Episode:    "Episode",
	Utterances: "Utterances",
}

// turnR is where relationships are stored.
type turnR struct {
	Episode    *Episode       `boil:"Episode" json:"Episode" toml:"Episode" yaml:"Episode"`
	Utterances UtteranceSlice `boil:"Utterances" json:"Utterances" toml:"Utterances" yaml:"Utterances"`
}

// NewStruct creates a new relationship struct
func (*turnR) NewStruct() *turnR {
	return &turnR{}
}

func (r *turnR) GetEpisode() *Episode {
	if r == nil {
		return nil
	}
	return r.Episode
}

func (r *turnR) GetUtterances() UtteranceSlice {
	if r == nil {
		return nil
	}
	return r.Utterances
}

// turnL is where Load methods for each relationship are stored.
type turnL struct{}

var (
	turnAllColumns            = []string{"id", "episode_id", "sequence_no", "start_time", "end_time", "created_at", "updated_at"}
	turnColumnsWithoutDefault = []string{"id", "episode_id", "sequence_no", "created_at", "updated_at"}
	turnColumnsWithDefault    = []string{"start_time", "end_time"}
	turnPrimaryKeyColumns     = []string{"id"}
	turnGeneratedColumns      = []string{}
)

type (
	// TurnSlice is an alias for a slice of pointers to Turn.
	// This should almost always be used instead of []Turn.
	TurnSlice []*Turn
	// TurnHook is the signature for custom Turn hook methods
	TurnHook func(context.Context, boil.ContextExecutor, *Turn) error

	turnQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	turnType                 = reflect.TypeOf(&Turn{})
	turnMapping              = queries.MakeStructMapping(turnType)
	turnPrimaryKeyMapping, _ = queries.BindMapping(turnType, turnMapping, turnPrimaryKeyColumns)
	turnInsertCacheMut       sync.RWMutex
	turnInsertCache          = make(map[string]insertCache)
	turnUpdateCacheMut       sync.RWMutex
	turnUpdateCache          = make(map[string]updateCache)
	turnUpsertCacheMut       sync.RWMutex
	turnUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var turnAfterSelectHooks []TurnHook

var turnBeforeInsertHooks []TurnHook
var turnAfterInsertHooks []TurnHook

var turnBeforeUpdateHooks []TurnHook
var turnAfterUpdateHooks []TurnHook

var turnBeforeDeleteHooks []TurnHook
var turnAfterDeleteHooks []TurnHook

var turnBeforeUpsertHooks []TurnHook
var turnAfterUpsertHooks []TurnHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Turn) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Turn) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Turn) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Turn) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Turn) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Turn) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Turn) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Turn) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Turn) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range turnAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTurnHook registers your hook function for all future operations.
func AddTurnHook(hookPoint boil.HookPoint, turnHook TurnHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		turnAfterSelectHooks = append(turnAfterSelectHooks, turnHook)
	case boil.BeforeInsertHook:
		turnBeforeInsertHooks = append(turnBeforeInsertHooks, turnHook)
	case boil.AfterInsertHook:
		turnAfterInsertHooks = append(turnAfterInsertHooks, turnHook)
	case boil.BeforeUpdateHook:
		turnBeforeUpdateHooks = append(turnBeforeUpdateHooks, turnHook)
	case boil.AfterUpdateHook:
		turnAfterUpdateHooks = append(turnAfterUpdateHooks, turnHook)
	case boil.BeforeDeleteHook:
		turnBeforeDeleteHooks = append(turnBeforeDeleteHooks, turnHook)
	case boil.AfterDeleteHook:
		turnAfterDeleteHooks = append(turnAfterDeleteHooks, turnHook)
	case boil.BeforeUpsertHook:
		turnBeforeUpsertHooks = append(turnBeforeUpsertHooks, turnHook)
	case boil.AfterUpsertHook:
		turnAfterUpsertHooks = append(turnAfterUpsertHooks, turnHook)
	}
}

// One returns a single turn record from the query.
func (q turnQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Turn, error) {
	o := &Turn{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datasource_raw: failed to execute a one query for turns")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Turn records from the query.
func (q turnQuery) All(ctx context.Context, exec boil.ContextExecutor) (TurnSlice, error) {
	var o []*Turn

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "datasource_raw: failed to assign all query results to Turn slice")
	}

	if len(turnAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Turn records in the query.
func (q turnQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to count turns rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q turnQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "datasource_raw: failed to check if turns exists")
	}

	return count > 0, nil
}

// Episode pointed to by the foreign key.
func (o *Turn) Episode(mods ...qm.QueryMod) episodeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EpisodeID),
	}

	queryMods = append(queryMods, mods...)

	return Episodes(queryMods...)
}

// Utterances retrieves all the utterance's Utterances with an executor.
func (o *Turn) Utterances(mods ...qm.QueryMod) utteranceQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"utterances\".\"turn_id\"=?", o.ID),
	)

	return Utterances(queryMods...)
}

// LoadEpisode allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (turnL) LoadEpisode(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTurn interface{}, mods queries.Applicator) error {
	var slice []*Turn
	var object *Turn

	if singular {
		var ok bool
		object, ok = maybeTurn.(*Turn)
		if !ok {
			object = new(Turn)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTurn)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTurn))
			}
		}
	} else {
		s, ok := maybeTurn.(*[]*Turn)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTurn)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTurn))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &turnR{}
		}
		args = append(args, object.EpisodeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &turnR{}
			}

			for _, a := range args {
				if a == obj.EpisodeID {
					continue Outer
				}
			}

			args = append(args, obj.EpisodeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`episodes`),
		qm.WhereIn(`episodes.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Episode")
	}

	var resultSlice []*Episode
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Episode")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for episodes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for episodes")
	}

	if len(turnAfterSelectHooks) != 0 {
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
		object.R.Episode = foreign
		if foreign.R == nil {
			foreign.R = &episodeR{}
		}
		foreign.R.Turns = append(foreign.R.Turns, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EpisodeID == foreign.ID {
				local.R.Episode = foreign
				if foreign.R == nil {
					foreign.R = &episodeR{}
				}
				foreign.R.Turns = append(foreign.R.Turns, local)
				break
			}
		}
	}

	return nil
}

// LoadUtterances allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (turnL) LoadUtterances(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTurn interface{}, mods queries.Applicator) error {
	var slice []*Turn
	var object *Turn

	if singular {
		var ok bool
		object, ok = maybeTurn.(*Turn)
		if !ok {
			object = new(Turn)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTurn)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTurn))
			}
		}
	} else {
		s, ok := maybeTurn.(*[]*Turn)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTurn)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTurn))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &turnR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &turnR{}
			}

			for _, a := range args {
				if a == obj.ID {
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
		qm.From(`utterances`),
		qm.WhereIn(`utterances.turn_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load utterances")
	}

	var resultSlice []*Utterance
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice utterances")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on utterances")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for utterances")
	}

	if len(utteranceAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Utterances = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &utteranceR{}
			}
			foreign.R.Turn = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TurnID {
				local.R.Utterances = append(local.R.Utterances, foreign)
				if foreign.R == nil {
					foreign.R = &utteranceR{}
				}
				foreign.R.Turn = local
				break
			}
		}
	}

	return nil
}

// SetEpisode of the turn to the related item.
// Sets o.R.Episode to related.
// Adds o to related.R.Turns.
func (o *Turn) SetEpisode(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Episode) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"turns\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"episode_id"}),
		strmangle.WhereClause("\"", "\"", 2, turnPrimaryKeyColumns),
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

	o.EpisodeID = related.ID
	if o.R == nil {
		o.R = &turnR{
			Episode: related,
		}
	} else {
		o.R.Episode = related
	}

	if related.R == nil {
		related.R = &episodeR{
			Turns: TurnSlice{o},
		}
	} else {
		related.R.Turns = append(related.R.Turns, o)
	}

	return nil
}

// AddUtterances adds the given related objects to the existing relationships
// of the turn, optionally inserting them as new records.
// Appends related to o.R.Utterances.
// Sets related.R.Turn appropriately.
func (o *Turn) AddUtterances(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Utterance) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TurnID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"utterances\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"turn_id"}),
				strmangle.WhereClause("\"", "\"", 2, utterancePrimaryKeyColumns),
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

			rel.TurnID = o.ID
		}
	}

	if o.R == nil {
		o.R = &turnR{
			Utterances: related,
		}
	} else {
		o.R.Utterances = append(o.R.Utterances, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &utteranceR{
				Turn: o,
			}
		} else {
			rel.R.Turn = o
		}
	}
	return nil
}

// Turns retrieves all the records using an executor.
func Turns(mods ...qm.QueryMod) turnQuery {
	mods = append(mods, qm.From("\"turns\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"turns\".*"})
	}

	return turnQuery{q}
}

// FindTurn retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTurn(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Turn, error) {
	turnObj := &Turn{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"turns\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, turnObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datasource_raw: unable to select from turns")
	}

	if err = turnObj.doAfterSelectHooks(ctx, exec); err != nil {
		return turnObj, err
	}

	return turnObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Turn) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("datasource_raw: no turns provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(turnColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	turnInsertCacheMut.RLock()
	cache, cached := turnInsertCache[key]
	turnInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			turnAllColumns,
			turnColumnsWithDefault,
			turnColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(turnType, turnMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(turnType, turnMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"turns\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"turns\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "datasource_raw: unable to insert into turns")
	}

	if !cached {
		turnInsertCacheMut.Lock()
		turnInsertCache[key] = cache
		turnInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Turn.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Turn) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	turnUpdateCacheMut.RLock()
	cache, cached := turnUpdateCache[key]
	turnUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			turnAllColumns,
			turnPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("datasource_raw: unable to update turns, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"turns\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, turnPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(turnType, turnMapping, append(wl, turnPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "datasource_raw: unable to update turns row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by update for turns")
	}

	if !cached {
		turnUpdateCacheMut.Lock()
		turnUpdateCache[key] = cache
		turnUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q turnQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to update all for turns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to retrieve rows affected for turns")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TurnSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), turnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"turns\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, turnPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to update all in turn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to retrieve rows affected all in update all turn")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Turn) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("datasource_raw: no turns provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(turnColumnsWithDefault, o)

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

	turnUpsertCacheMut.RLock()
	cache, cached := turnUpsertCache[key]
	turnUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			turnAllColumns,
			turnColumnsWithDefault,
			turnColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			turnAllColumns,
			turnPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("datasource_raw: unable to upsert turns, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(turnPrimaryKeyColumns))
			copy(conflict, turnPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"turns\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(turnType, turnMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(turnType, turnMapping, ret)
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
		return errors.Wrap(err, "datasource_raw: unable to upsert turns")
	}

	if !cached {
		turnUpsertCacheMut.Lock()
		turnUpsertCache[key] = cache
		turnUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Turn record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Turn) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("datasource_raw: no Turn provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), turnPrimaryKeyMapping)
	sql := "DELETE FROM \"turns\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to delete from turns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by delete for turns")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q turnQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("datasource_raw: no turnQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to delete all from turns")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by deleteall for turns")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TurnSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(turnBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), turnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"turns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, turnPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: unable to delete all from turn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datasource_raw: failed to get rows affected by deleteall for turns")
	}

	if len(turnAfterDeleteHooks) != 0 {
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
func (o *Turn) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTurn(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TurnSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TurnSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), turnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"turns\".* FROM \"turns\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, turnPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "datasource_raw: unable to reload all in TurnSlice")
	}

	*o = slice

	return nil
}

// TurnExists checks if the Turn row exists.
func TurnExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"turns\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "datasource_raw: unable to check if turns exists")
	}

	return exists, nil
}
