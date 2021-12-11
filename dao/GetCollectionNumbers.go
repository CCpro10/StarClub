package dao

import (
	"StarClub/model"
	"github.com/go-redis/redis/v8"
	"strconv"
)

//给活动切片添加上收藏人数
func GetActivitiesCollectionNumbers(activities []model.Activity) (err error) {
	for k, _ := range activities {
		nums, err := RedisDB.Get(CTX, strconv.Itoa(int(activities[k].ID))).Result()
		//没有查到，继续查下一个
		if err == redis.Nil {
			continue
			//查到了
		} else {
			if activities[k].CollectionNumbers, err = strconv.Atoi(nums); err != nil {
				return err
			}
		}
	}
	return nil
}

//获取某个活动的收藏人数
func GetActivityCollectionNumbers(activity *model.Activity) (err error) {
	nums, err := RedisDB.Get(CTX, strconv.Itoa(int(activity.ID))).Result()
	//没有查到
	if err == redis.Nil {
		return nil
	} else if err != nil { //查询错误
		return err
	} else { //查到了
		if activity.CollectionNumbers, err = strconv.Atoi(nums); err != nil {
			return err
		}
	}
	return nil
}
