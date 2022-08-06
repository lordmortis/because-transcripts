// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datasource_raw

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Episodes", testEpisodes)
	t.Run("Podcasts", testPodcasts)
	t.Run("Speakers", testSpeakers)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinks)
	t.Run("UtteranceFragments", testUtteranceFragments)
	t.Run("Utterances", testUtterances)
}

func TestDelete(t *testing.T) {
	t.Run("Episodes", testEpisodesDelete)
	t.Run("Podcasts", testPodcastsDelete)
	t.Run("Speakers", testSpeakersDelete)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksDelete)
	t.Run("UtteranceFragments", testUtteranceFragmentsDelete)
	t.Run("Utterances", testUtterancesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Episodes", testEpisodesQueryDeleteAll)
	t.Run("Podcasts", testPodcastsQueryDeleteAll)
	t.Run("Speakers", testSpeakersQueryDeleteAll)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksQueryDeleteAll)
	t.Run("UtteranceFragments", testUtteranceFragmentsQueryDeleteAll)
	t.Run("Utterances", testUtterancesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Episodes", testEpisodesSliceDeleteAll)
	t.Run("Podcasts", testPodcastsSliceDeleteAll)
	t.Run("Speakers", testSpeakersSliceDeleteAll)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksSliceDeleteAll)
	t.Run("UtteranceFragments", testUtteranceFragmentsSliceDeleteAll)
	t.Run("Utterances", testUtterancesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Episodes", testEpisodesExists)
	t.Run("Podcasts", testPodcastsExists)
	t.Run("Speakers", testSpeakersExists)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksExists)
	t.Run("UtteranceFragments", testUtteranceFragmentsExists)
	t.Run("Utterances", testUtterancesExists)
}

func TestFind(t *testing.T) {
	t.Run("Episodes", testEpisodesFind)
	t.Run("Podcasts", testPodcastsFind)
	t.Run("Speakers", testSpeakersFind)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksFind)
	t.Run("UtteranceFragments", testUtteranceFragmentsFind)
	t.Run("Utterances", testUtterancesFind)
}

func TestBind(t *testing.T) {
	t.Run("Episodes", testEpisodesBind)
	t.Run("Podcasts", testPodcastsBind)
	t.Run("Speakers", testSpeakersBind)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksBind)
	t.Run("UtteranceFragments", testUtteranceFragmentsBind)
	t.Run("Utterances", testUtterancesBind)
}

func TestOne(t *testing.T) {
	t.Run("Episodes", testEpisodesOne)
	t.Run("Podcasts", testPodcastsOne)
	t.Run("Speakers", testSpeakersOne)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksOne)
	t.Run("UtteranceFragments", testUtteranceFragmentsOne)
	t.Run("Utterances", testUtterancesOne)
}

func TestAll(t *testing.T) {
	t.Run("Episodes", testEpisodesAll)
	t.Run("Podcasts", testPodcastsAll)
	t.Run("Speakers", testSpeakersAll)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksAll)
	t.Run("UtteranceFragments", testUtteranceFragmentsAll)
	t.Run("Utterances", testUtterancesAll)
}

func TestCount(t *testing.T) {
	t.Run("Episodes", testEpisodesCount)
	t.Run("Podcasts", testPodcastsCount)
	t.Run("Speakers", testSpeakersCount)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksCount)
	t.Run("UtteranceFragments", testUtteranceFragmentsCount)
	t.Run("Utterances", testUtterancesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Episodes", testEpisodesHooks)
	t.Run("Podcasts", testPodcastsHooks)
	t.Run("Speakers", testSpeakersHooks)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksHooks)
	t.Run("UtteranceFragments", testUtteranceFragmentsHooks)
	t.Run("Utterances", testUtterancesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Episodes", testEpisodesInsert)
	t.Run("Episodes", testEpisodesInsertWhitelist)
	t.Run("Podcasts", testPodcastsInsert)
	t.Run("Podcasts", testPodcastsInsertWhitelist)
	t.Run("Speakers", testSpeakersInsert)
	t.Run("Speakers", testSpeakersInsertWhitelist)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksInsert)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksInsertWhitelist)
	t.Run("UtteranceFragments", testUtteranceFragmentsInsert)
	t.Run("UtteranceFragments", testUtteranceFragmentsInsertWhitelist)
	t.Run("Utterances", testUtterancesInsert)
	t.Run("Utterances", testUtterancesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("EpisodeToPodcastUsingPodcast", testEpisodeToOnePodcastUsingPodcast)
	t.Run("UtteranceFragmentLinkToUtteranceFragmentUsingUtteranceFragment", testUtteranceFragmentLinkToOneUtteranceFragmentUsingUtteranceFragment)
	t.Run("UtteranceFragmentLinkToUtteranceUsingUtterance", testUtteranceFragmentLinkToOneUtteranceUsingUtterance)
	t.Run("UtteranceToEpisodeUsingEpisode", testUtteranceToOneEpisodeUsingEpisode)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("EpisodeToUtterances", testEpisodeToManyUtterances)
	t.Run("PodcastToEpisodes", testPodcastToManyEpisodes)
	t.Run("SpeakerToUtterances", testSpeakerToManyUtterances)
	t.Run("UtteranceFragmentToUtteranceFragmentLinks", testUtteranceFragmentToManyUtteranceFragmentLinks)
	t.Run("UtteranceToUtteranceFragmentLinks", testUtteranceToManyUtteranceFragmentLinks)
	t.Run("UtteranceToSpeakers", testUtteranceToManySpeakers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("EpisodeToPodcastUsingEpisodes", testEpisodeToOneSetOpPodcastUsingPodcast)
	t.Run("UtteranceFragmentLinkToUtteranceFragmentUsingUtteranceFragmentLinks", testUtteranceFragmentLinkToOneSetOpUtteranceFragmentUsingUtteranceFragment)
	t.Run("UtteranceFragmentLinkToUtteranceUsingUtteranceFragmentLinks", testUtteranceFragmentLinkToOneSetOpUtteranceUsingUtterance)
	t.Run("UtteranceToEpisodeUsingUtterances", testUtteranceToOneSetOpEpisodeUsingEpisode)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("EpisodeToUtterances", testEpisodeToManyAddOpUtterances)
	t.Run("PodcastToEpisodes", testPodcastToManyAddOpEpisodes)
	t.Run("SpeakerToUtterances", testSpeakerToManyAddOpUtterances)
	t.Run("UtteranceFragmentToUtteranceFragmentLinks", testUtteranceFragmentToManyAddOpUtteranceFragmentLinks)
	t.Run("UtteranceToUtteranceFragmentLinks", testUtteranceToManyAddOpUtteranceFragmentLinks)
	t.Run("UtteranceToSpeakers", testUtteranceToManyAddOpSpeakers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("SpeakerToUtterances", testSpeakerToManySetOpUtterances)
	t.Run("UtteranceToSpeakers", testUtteranceToManySetOpSpeakers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("SpeakerToUtterances", testSpeakerToManyRemoveOpUtterances)
	t.Run("UtteranceToSpeakers", testUtteranceToManyRemoveOpSpeakers)
}

func TestReload(t *testing.T) {
	t.Run("Episodes", testEpisodesReload)
	t.Run("Podcasts", testPodcastsReload)
	t.Run("Speakers", testSpeakersReload)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksReload)
	t.Run("UtteranceFragments", testUtteranceFragmentsReload)
	t.Run("Utterances", testUtterancesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Episodes", testEpisodesReloadAll)
	t.Run("Podcasts", testPodcastsReloadAll)
	t.Run("Speakers", testSpeakersReloadAll)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksReloadAll)
	t.Run("UtteranceFragments", testUtteranceFragmentsReloadAll)
	t.Run("Utterances", testUtterancesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Episodes", testEpisodesSelect)
	t.Run("Podcasts", testPodcastsSelect)
	t.Run("Speakers", testSpeakersSelect)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksSelect)
	t.Run("UtteranceFragments", testUtteranceFragmentsSelect)
	t.Run("Utterances", testUtterancesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Episodes", testEpisodesUpdate)
	t.Run("Podcasts", testPodcastsUpdate)
	t.Run("Speakers", testSpeakersUpdate)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksUpdate)
	t.Run("UtteranceFragments", testUtteranceFragmentsUpdate)
	t.Run("Utterances", testUtterancesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Episodes", testEpisodesSliceUpdateAll)
	t.Run("Podcasts", testPodcastsSliceUpdateAll)
	t.Run("Speakers", testSpeakersSliceUpdateAll)
	t.Run("UtteranceFragmentLinks", testUtteranceFragmentLinksSliceUpdateAll)
	t.Run("UtteranceFragments", testUtteranceFragmentsSliceUpdateAll)
	t.Run("Utterances", testUtterancesSliceUpdateAll)
}
