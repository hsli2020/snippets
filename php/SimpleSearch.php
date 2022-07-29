<?php

use Illuminate\Database\Eloquent\Builder;

/**
 * Usage:
 * protected $searchable = [
 *     'title:3',
 *     'content:1',
 * ];
 *
 * $this->simpleSearch($phrase);
 *
 * @method static \Illuminate\Database\Eloquent\Builder simpleSearch($phrase, $matchAll, $withoutScore)
 */
trait SimpleSearch
{
    /**
     * Scope a query to only filter result based on a simple phrase search,
     * with optional relevance score.
     *
     * @param \Illuminate\Database\Eloquent\Builder $query
     * @param string $phrase
     * @param boolean $matchAll
     * @param boolean $withoutScore
     *
     * @return \Illuminate\Database\Eloquent\Builder
     */
    public function scopeSimpleSearch(Builder $query, string $phrase, $matchAll = false, $withoutScore = false)
    {
        // Strip double spaces and explode to keywords array
        $keywords = explode(" ", preg_replace("/\s+/", " ", $phrase));

        // Add relevance score to select
        if ($withoutScore) {
            $query
                ->select('*')
                ->addSelect($this->getScoreSelect($keywords));
        }

        // Search with an OR WHERE LIKE query
        foreach ($this->getSearchableColumns() as $column => $weight) {
            $query->where(function ($query) use ($keywords, $column) {
                foreach ($keywords as $keyword) {
                    $query->where($column, 'LIKE', "%{$keyword}%", $matchAll ? 'and' : 'or');
                }
            });
        }

        return $query;
    }

    /**
     * Small algorithm to add search relevance to select
     *
     * @param string $phrase
     *
     * @return string
     */
    private function getScoreSelect($keywords): string
    {
        $selects = [];
        $totalWeight = 0;

        foreach ($this->getSearchableColumns() as $column => $weight) {
            $totalWeight += $weight;

            $selects[] = '(
                (
                    CHAR_LENGTH(' . $column . ') -
                    CHAR_LENGTH(REGEXP_REPLACE(' . $column . ', "(' . implode('|', $keywords) . ')", ""))
                ) / SQRT(CHAR_LENGTH(' . $column . '))
            ) * ' . intval($weight);
        }

        return '(' . implode(' + ', $selects) . ') / ' . $totalWeight . ' as score';
    }

    /**
     * Get searchable columns as defined in class
     *
     * @return array
     */
    private function getSearchableColumns(): array
    {
        $columns = [];

        foreach ($this->searchable ?? [] as $column) {
            $columnData = explode(':', $column);
            $columnName = $columnData[0];
            $columnWeight = $columnData[1] ?? 1;

            $columns[$columnName] = $columnWeight;
        }

        return $columns;
    }
}
