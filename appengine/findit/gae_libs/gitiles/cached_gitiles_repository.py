# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from gae_libs.caches import CompressedMemCache
from gae_libs.caches import PickledMemCache

from lib.cache_decorator import Cached
from lib.gitiles.gitiles_repository import GitilesRepository

CACHE_EXPIRE_TIME_SECONDS = 24 * 60 * 60


class CachedGitilesRepository(GitilesRepository):

  @Cached(namespace='Gitiles-json-view', expire_time=CACHE_EXPIRE_TIME_SECONDS,
          cache=CompressedMemCache())
  def _SendRequestForJsonResponse(self, url, params=None):  # pragma: no cover
    return super(CachedGitilesRepository, self)._SendRequestForJsonResponse(
        url, params=params)

  @Cached(namespace='Gitiles-text-view', expire_time=CACHE_EXPIRE_TIME_SECONDS,
          cache=PickledMemCache())
  def _SendRequestForTextResponse(self, url):  # pragma: no cover
    return super(CachedGitilesRepository, self)._SendRequestForTextResponse(url)
