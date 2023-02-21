package core

import (
	"context"
	//"strconv"
	//"strings"
	//"study/TikTok/config"
	//"study/TikTok/middleware/rabbitmq"
	//"study/TikTok/middleware/redis"
	//"time"
	//"video_service/dao"
	"video_service/services"
)

//视频流
func (*VideoService) FeedInfo(ctx context.Context, req *services.FeedRequest, resp *services.FeedResponse) error {
	resp.StatusCode = 1
	resp.StatusMsg = "okk"

	return nil
}

//发布视频
func (*VideoService) PublishAction(ctx context.Context, req *services.PublishActionRequest, resp *services.PublishActionResponse) error {
	return nil
}

//自己的发布视频列表
func (*VideoService) PublishList(ctx context.Context, req *services.PublishListRequest, resp *services.PublishListResponse) error {
	return nil
}

//点赞/取消点赞
func (*VideoService) FavoriteAction(ctx context.Context, req *services.FavoriteActionRequest, resp *services.FavoriteActionResponse) error {
	//strUserId := ctx.GetString()
	////将int64 videoId转换为 string strVideoId
	//strUserId := strconv.FormatInt(userId, 10)
	////将int64 videoId转换为 string strVideoId
	//strVideoId := strconv.FormatInt(videoId, 10)
	////将要操作数据库likes表的信息打入消息队列RmqLikeAdd或者RmqLikeDel
	////拼接打入信息
	//sb := strings.Builder{}
	//sb.WriteString(strUserId)
	//sb.WriteString(" ")
	//sb.WriteString(strVideoId)
	//
	////step1:维护Redis LikeUserId、LikeVideoId;
	////执行点赞操作维护
	//if actionType == config.LikeAction {
	//	//查询Redis LikeUserId(key:strUserId)是否已经加载过此信息
	//	if n, err := redis.RdbLikeUserId.Exists(redis.Ctx, strUserId).Result(); n > 0 {
	//		//如果有问题，说明查询redis失败,返回错误信息
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId query key失败：%v", err)
	//			return err
	//		} //如果加载过此信息key:strUserId，则加入value:videoId
	//		//如果redis LikeUserId 添加失败，数据库操作成功，会有脏数据，所以只有redis操作成功才执行数据库likes表操作
	//		if _, err1 := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, videoId).Result(); err1 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId add value失败：%v", err1)
	//			return err1
	//		} else {
	//			//如果数据库操作失败了，redis是正确数据，客户端显示的是点赞成功，不会影响后续结果
	//			//只有当该用户取消所有点赞视频的时候redis才会重新加载数据库信息，这时候因为取消赞了必然和数据库信息一致
	//			//同样这条信息消费成功与否也不重要，因为redis是正确信息,理由如上
	//			rabbitmq.RmqLikeAdd.Publish(sb.String())
	//		}
	//	} else { //如果不存在，则维护Redis LikeUserId 新建key:strUserId,设置过期时间，加入DefaultRedisValue，
	//		//key:strUserId，加入value:DefaultRedisValue,过期才会删，防止删最后一个数据的时候数据库还没更新完出现脏读，或者数据库操作失败造成的脏读
	//		//通过userId查询likes表,返回所有点赞videoId，加入key:strUserId集合中,
	//		//再加入当前videoId,再更新likes表此条数据
	//		if _, err := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, config.DefaultRedisValue).Result(); err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId add value失败")
	//			redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
	//			return err
	//		}
	//		//给键值设置有效期，类似于gc机制
	//		_, err := redis.RdbLikeUserId.Expire(redis.Ctx, strUserId,
	//			time.Duration(config.OneMonth)*time.Second).Result()
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId 设置有效期失败")
	//			redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
	//			return err
	//		}
	//		videoIdList, err1 := dao.GetLikeVideoIdList(userId)
	//		//如果有问题，说明查询失败，返回错误信息："get likeVideoIdList failed"
	//		if err1 != nil {
	//			return err1
	//		}
	//		//遍历videoIdList,添加进key的集合中，若失败，删除key，并返回错误信息，这么做的原因是防止脏读，
	//		//保证redis与mysql数据一致性
	//		for _, likeVideoId := range videoIdList {
	//			if _, err1 := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, likeVideoId).Result(); err1 != nil {
	//				log.Printf("方法:FavouriteAction RedisLikeUserId add value失败")
	//				redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
	//				return err1
	//			}
	//		}
	//		//这样操作理由同上
	//		if _, err2 := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, videoId).Result(); err2 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId add value失败：%v", err2)
	//			return err2
	//		} else {
	//			rabbitmq.RmqLikeAdd.Publish(sb.String())
	//		}
	//	}
	//
	//	//查询Redis LikeVideoId(key:strVideoId)是否已经加载过此信息
	//	if n, err := redis.RdbLikeVideoId.Exists(redis.Ctx, strVideoId).Result(); n > 0 {
	//		//如果有问题，说明查询redis失败,返回错误信息
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId query key失败：%v", err)
	//			return err
	//		} //如果加载过此信息key:strVideoId，则加入value:userId
	//		//如果redis LikeVideoId 添加失败，返回错误信息
	//		if _, err1 := redis.RdbLikeVideoId.SAdd(redis.Ctx, strVideoId, userId).Result(); err1 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId add value失败：%v", err1)
	//			return err1
	//		}
	//	} else { //如果不存在，则维护Redis LikeVideoId 新建key:strVideoId，设置有效期，加入DefaultRedisValue
	//		//通过videoId查询likes表,返回所有点赞userId，加入key:strVideoId集合中,
	//		//再加入当前userId,再更新likes表此条数据
	//		//key:strVideoId，加入value:DefaultRedisValue,过期才会删，防止删最后一个数据的时候数据库还没更新完出现脏读，或者数据库操作失败造成的脏读
	//		if _, err := redis.RdbLikeVideoId.SAdd(redis.Ctx, strVideoId, config.DefaultRedisValue).Result(); err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId add value失败")
	//			redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//			return err
	//		}
	//		//给键值设置有效期，类似于gc机制
	//		_, err := redis.RdbLikeVideoId.Expire(redis.Ctx, strVideoId,
	//			time.Duration(config.OneMonth)*time.Second).Result()
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId 设置有效期失败")
	//			redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//			return err
	//		}
	//		userIdList, err1 := dao.GetLikeUserIdList(videoId)
	//		//如果有问题，说明查询失败，返回错误信息："get likeUserIdList failed"
	//		if err1 != nil {
	//			return err1
	//		}
	//		//遍历userIdList,添加进key的集合中，若失败，删除key，并返回错误信息，这么做的原因是防止脏读，
	//		//保证redis与mysql数据一致性
	//		for _, likeUserId := range userIdList {
	//			if _, err1 := redis.RdbLikeVideoId.SAdd(redis.Ctx, strVideoId, likeUserId).Result(); err1 != nil {
	//				log.Printf("方法:FavouriteAction RedisLikeVideoId add value失败")
	//				redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//				return err1
	//			}
	//		}
	//		//这样操作理由同上
	//		if _, err2 := redis.RdbLikeVideoId.SAdd(redis.Ctx, strVideoId, userId).Result(); err2 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId add value失败：%v", err2)
	//			return err2
	//		}
	//	}
	//} else { //执行取消赞操作维护
	//	//查询Redis LikeUserId(key:strUserId)是否已经加载过此信息
	//	if n, err := redis.RdbLikeUserId.Exists(redis.Ctx, strUserId).Result(); n > 0 {
	//		//如果有问题，说明查询redis失败,返回错误信息
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId query key失败：%v", err)
	//			return err
	//		} //防止出现redis数据不一致情况，当redis删除操作成功，才执行数据库更新操作
	//		if _, err1 := redis.RdbLikeUserId.SRem(redis.Ctx, strUserId, videoId).Result(); err1 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId del value失败：%v", err1)
	//			return err1
	//		} else {
	//			//后续数据库的操作，可以在mq里设置若执行数据库更新操作失败，重新消费该信息
	//			rabbitmq.RmqLikeDel.Publish(sb.String())
	//		}
	//	} else { //如果不存在，则维护Redis LikeUserId 新建key:strUserId，加入value:DefaultRedisValue,过期才会删，防止删最后一个数据的时候数据库
	//		// 还没更新完出现脏读，或者数据库操作失败造成的脏读
	//		//通过userId查询likes表,返回所有点赞videoId，加入key:strUserId集合中,
	//		//再删除当前videoId,再更新likes表此条数据
	//		//key:strUserId，加入value:DefaultRedisValue,过期才会删，防止删最后一个数据的时候数据库还没更新完出现脏读，或者数据库操作失败造成的脏读
	//		if _, err := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, config.DefaultRedisValue).Result(); err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId add value失败")
	//			redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
	//			return err
	//		}
	//		//给键值设置有效期，类似于gc机制
	//		_, err := redis.RdbLikeUserId.Expire(redis.Ctx, strUserId,
	//			time.Duration(config.OneMonth)*time.Second).Result()
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId 设置有效期失败")
	//			redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
	//			return err
	//		}
	//		videoIdList, err1 := dao.GetLikeVideoIdList(userId)
	//		//如果有问题，说明查询失败，返回错误信息："get likeVideoIdList failed"
	//		if err1 != nil {
	//			return err1
	//		}
	//		//遍历videoIdList,添加进key的集合中，若失败，删除key，并返回错误信息，这么做的原因是防止脏读，
	//		//保证redis与mysql 数据原子性
	//		for _, likeVideoId := range videoIdList {
	//			if _, err1 := redis.RdbLikeUserId.SAdd(redis.Ctx, strUserId, likeVideoId).Result(); err1 != nil {
	//				log.Printf("方法:FavouriteAction RedisLikeUserId add value失败")
	//				redis.RdbLikeUserId.Del(redis.Ctx, strUserId)
	//				return err1
	//			}
	//		}
	//		//这样操作理由同上
	//		if _, err2 := redis.RdbLikeUserId.SRem(redis.Ctx, strUserId, videoId).Result(); err2 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeUserId del value失败：%v", err2)
	//			return err2
	//		} else {
	//			rabbitmq.RmqLikeDel.Publish(sb.String())
	//		}
	//	}
	//
	//	//查询Redis LikeVideoId(key:strVideoId)是否已经加载过此信息
	//	if n, err := redis.RdbLikeVideoId.Exists(redis.Ctx, strVideoId).Result(); n > 0 {
	//		//如果有问题，说明查询redis失败,返回错误信息
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId query key失败：%v", err)
	//			return err
	//		} //如果加载过此信息key:strVideoId，则删除value:userId
	//		//如果redis LikeVideoId 删除失败，返回错误信息
	//		if _, err1 := redis.RdbLikeVideoId.SRem(redis.Ctx, strVideoId, userId).Result(); err1 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId del value失败：%v", err1)
	//			return err1
	//		}
	//	} else { //如果不存在，则维护Redis LikeVideoId 新建key:strVideoId,加入value:DefaultRedisValue,过期才会删，防止删最后一个数据的时候数据库
	//		// 还没更新完出现脏读，或者数据库操作失败造成的脏读
	//		//通过videoId查询likes表,返回所有点赞userId，加入key:strVideoId集合中,
	//		//再删除当前userId,再更新likes表此条数据
	//		//key:strVideoId，加入value:DefaultRedisValue,过期才会删，防止删最后一个数据的时候数据库还没更新完出现脏读，或者数据库操作失败造成的脏读
	//		if _, err := redis.RdbLikeVideoId.SAdd(redis.Ctx, strVideoId, config.DefaultRedisValue).Result(); err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId add value失败")
	//			redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//			return err
	//		}
	//		//给键值设置有效期，类似于gc机制
	//		_, err := redis.RdbLikeVideoId.Expire(redis.Ctx, strVideoId,
	//			time.Duration(config.OneMonth)*time.Second).Result()
	//		if err != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId 设置有效期失败")
	//			redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//			return err
	//		}
	//
	//		userIdList, err1 := dao.GetLikeUserIdList(videoId)
	//		//如果有问题，说明查询失败，返回错误信息："get likeUserIdList failed"
	//		if err1 != nil {
	//			redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//			return err1
	//		}
	//		//遍历userIdList,添加进key的集合中，若失败，删除key，并返回错误信息，这么做的原因是防止脏读，
	//		//保证redis与mysql数据一致性
	//		for _, likeUserId := range userIdList {
	//			if _, err1 := redis.RdbLikeVideoId.SAdd(redis.Ctx, strVideoId, likeUserId).Result(); err1 != nil {
	//				log.Printf("方法:FavouriteAction RedisLikeVideoId add value失败")
	//				redis.RdbLikeVideoId.Del(redis.Ctx, strVideoId)
	//				return err1
	//			}
	//		}
	//		//这样操作理由同上
	//		if _, err2 := redis.RdbLikeVideoId.SRem(redis.Ctx, strVideoId, userId).Result(); err2 != nil {
	//			log.Printf("方法:FavouriteAction RedisLikeVideoId del value失败：%v", err2)
	//			return err2
	//		}
	//	}
	//}
	return nil
}

//点赞列表
func (*VideoService) FavoriteList(ctx context.Context, req *services.FavoriteListRequest, resp *services.FavoriteListResponse) error {
	return nil
}

//评论
func (*VideoService) CommentAction(ctx context.Context, req *services.CommentActionRequest, resp *services.CommentActionResponse) error {
	return nil
}

//评论列表
func (*VideoService) CommentList(ctx context.Context, req *services.CommentListRequest, resp *services.CommentListResponse) error {
	return nil
}
