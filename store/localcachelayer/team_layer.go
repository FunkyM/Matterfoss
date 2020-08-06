// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package localcachelayer

import (
	"github.com/cjdelisle/matterfoss-server/v5/model"
	"github.com/cjdelisle/matterfoss-server/v5/store"
)

type LocalCacheTeamStore struct {
	store.TeamStore
	rootStore *LocalCacheStore
}

func (s *LocalCacheTeamStore) handleClusterInvalidateTeam(msg *model.ClusterMessage) {
	if msg.Data == CLEAR_CACHE_MESSAGE_DATA {
		s.rootStore.teamAllTeamIdsForUserCache.Purge()
	} else {
		s.rootStore.teamAllTeamIdsForUserCache.Remove(msg.Data)
	}
}

func (s LocalCacheTeamStore) ClearCaches() {
	s.rootStore.teamAllTeamIdsForUserCache.Purge()
	if s.rootStore.metrics != nil {
		s.rootStore.metrics.IncrementMemCacheInvalidationCounter("All Team Ids for User - Purge")
	}
}

func (s LocalCacheTeamStore) InvalidateAllTeamIdsForUser(userId string) {
	s.rootStore.doInvalidateCacheCluster(s.rootStore.teamAllTeamIdsForUserCache, userId)
	if s.rootStore.metrics != nil {
		s.rootStore.metrics.IncrementMemCacheInvalidationCounter("All Team Ids for User - Remove by UserId")
	}
}

func (s LocalCacheTeamStore) GetUserTeamIds(userID string, allowFromCache bool) ([]string, *model.AppError) {
	if !allowFromCache {
		return s.TeamStore.GetUserTeamIds(userID, allowFromCache)
	}

	var userTeamIds []string
	if err := s.rootStore.doStandardReadCache(s.rootStore.teamAllTeamIdsForUserCache, userID, &userTeamIds); err == nil {
		return userTeamIds, nil
	}

	userTeamIds, err := s.TeamStore.GetUserTeamIds(userID, allowFromCache)
	if err != nil {
		return nil, err
	}

	if len(userTeamIds) > 0 {
		s.rootStore.doStandardAddToCache(s.rootStore.teamAllTeamIdsForUserCache, userID, userTeamIds)
	}

	return userTeamIds, nil
}

func (s LocalCacheTeamStore) Update(team *model.Team) (*model.Team, *model.AppError) {
	var oldTeam *model.Team
	var err *model.AppError
	if team.DeleteAt != 0 {
		oldTeam, err = s.TeamStore.Get(team.Id)
		if err != nil {
			return nil, err
		}
	}

	tm, err := s.TeamStore.Update(team)
	if err != nil {
		return nil, err
	}
	defer s.rootStore.doClearCacheCluster(s.rootStore.rolePermissionsCache)

	if oldTeam != nil && oldTeam.DeleteAt == 0 {
		s.rootStore.doClearCacheCluster(s.rootStore.teamAllTeamIdsForUserCache)
	}

	return tm, err
}
