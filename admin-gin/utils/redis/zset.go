package redis

import (
	"admin-gin/global"
	"context"
	"github.com/redis/go-redis/v9"
)

type ResZSet struct {
	Client redis.UniversalClient
}

func NewResZSet() *ResZSet {
	return &ResZSet{
		Client: global.GVA_REDIS,
	}
}

// 添加成员到有序集合
// 添加或更新有序集合中指定成员的分数
func (r *ResZSet) AddMember(ctx context.Context, key string, member string, score float64) error {
	z := &redis.Z{
		Score:  score,
		Member: member,
	}
	return r.Client.ZAdd(ctx, key, *z).Err()
}

// 获取有序集合中的成员，按照分数降序排列   -- 降序!!!
// start 和 stop 分别是排名的起始和结束位置，例如 0 表示第一个元素，-1 表示最后一个元素
func (r *ResZSet) GetMembers(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	//vals, err := r.Client.ZRangeWithScores(ctx, key, start, stop).Result()  升序
	vals, err := r.Client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return vals, nil
}

// 根据分数区间获取有序集合中的成员
// min 和 max 是分数的最小和最大值，负无穷和正无穷可以使用 "-inf" 和 "+inf"
func (r *ResZSet) GetMembersByScore(ctx context.Context, key string, min, max string) ([]redis.Z, error) {
	vals, err := r.Client.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()
	if err != nil {
		return nil, err
	}
	return vals, nil
}

// 删除有序集合中的某个成员
func (r *ResZSet) RemoveMember(ctx context.Context, key, member string) error {
	return r.Client.ZRem(ctx, key, member).Err()
}

// 获取成员的排名（按分数排序）
// 返回的是该成员在有序集合中的排名，排名从 0 开始
func (r *ResZSet) GetMemberRank(ctx context.Context, key, member string) (int64, error) {
	rank, err := r.Client.ZRank(ctx, key, member).Result()
	if err != nil {
		return 0, err
	}
	return rank, nil
}

// 获取有序集合成员的分数
func (r *ResZSet) GetMemberScore(ctx context.Context, key, member string) (float64, error) {
	score, err := r.Client.ZScore(ctx, key, member).Result()
	if err != nil {
		return 0, err
	}
	return score, nil
}

// 给某个成员增加指定分数
// increment 是增加的分数值，可以是负数或正数，1 表示加一分
func (r *ResZSet) IncrementMemberScore(ctx context.Context, key string, member string, increment float64) (float64, error) {
	newScore, err := r.Client.ZIncrBy(ctx, key, increment, member).Result()
	if err != nil {
		return 0, err
	}
	return newScore, nil
}
