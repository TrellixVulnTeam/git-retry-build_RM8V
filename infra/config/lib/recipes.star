# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Functions and constants related to recipes ecosystem support."""

load('//lib/infra.star', 'infra')


# Recipes module => name to use in builders.
_FRIENDLY_NAME = {
    'build': 'Build',
    'depot_tools': 'Depot Tools',
    'infra': 'Infra',
    'recipe_engine': 'Recipe Engine',
    'skia': 'Skia',
    'skiabuildbot': 'Skia Buildbot',
}


def _friendly(proj):
  return _FRIENDLY_NAME.get(proj, proj)


def _recipes():
  """Defines all recipes used by this module."""
  infra.recipe(name = 'recipe_simulation')
  infra.recipe(name = 'recipe_roll_tryjob')


def simulation_tester(
      name,
      project_under_test,
      triggered_by,
      console_view=None,
      console_category=None,
  ):
  """Defines a CI builder that runs recipe simulation tests."""
  luci.builder(
      name = name,
      bucket = 'ci',
      recipe = 'recipe_simulation',
      properties = {'project_under_test': project_under_test},
      dimensions = {
          'os': 'Ubuntu-14.04',
          'cpu': 'x86-64',
          'pool': 'luci.flex.ci',
      },
      service_account = infra.SERVICE_ACCOUNT_CI,
      build_numbers = True,
      execution_timeout = 30 * time.minute,
      swarming_tags = ['vpython:native-python-wrapper'],
      triggered_by = [triggered_by],
  )
  if console_view:
    luci.console_view_entry(
        builder = name,
        console_view = console_view,
        category = console_category,
    )


def roll_trybots(upstream, downstream, cq_group):
  """Defines a bunch of recipe roller trybots, one per downstream project."""
  for proj in downstream:
    name = '%s downstream Recipe Roll tester from %s' % (_friendly(proj), _friendly(upstream))
    luci.builder(
        name = name,
        bucket = 'try',
        recipe = 'recipe_roll_tryjob',
        properties = {
            'upstream_project': upstream,
            'downstream_project': proj,
        },
        dimensions = {
            'os': 'Ubuntu-14.04',
            'cpu': 'x86-64',
            'pool': 'luci.flex.try',
        },
        service_account = infra.SERVICE_ACCOUNT_TRY,
        execution_timeout = 30 * time.minute,
        swarming_tags = ['vpython:native-python-wrapper'],
    )
    luci.cq_tryjob_verifier(
        builder = name,
        cq_group = cq_group,
    )


recipes = struct(
    recipes = _recipes,
    simulation_tester = simulation_tester,
    roll_trybots = roll_trybots,
)
