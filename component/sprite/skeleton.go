package sprite

var (
	// Skeleton Death Sprites.
	SkeletonDeath1 = Graphic("skeleton_death_1")
	SkeletonDeath2 = Graphic("skeleton_death_2")
	SkeletonDeath3 = Graphic("skeleton_death_3")
	SkeletonDeath4 = Graphic("skeleton_death_4")
	SkeletonDeath5 = Graphic("skeleton_death_5")
	SkeletonDeath6 = Graphic("skeleton_death_6")
	// Skeleton Move Up Sprites.
	SkeletonMoveUp1 = Graphic("skeleton_move_up_1") // idle
	SkeletonMoveUp2 = Graphic("skeleton_move_up_2")
	SkeletonMoveUp3 = Graphic("skeleton_move_up_3")
	SkeletonMoveUp4 = Graphic("skeleton_move_up_4")
	SkeletonMoveUp5 = Graphic("skeleton_move_up_5")
	SkeletonMoveUp6 = Graphic("skeleton_move_up_6")
	SkeletonMoveUp7 = Graphic("skeleton_move_up_7")
	SkeletonMoveUp8 = Graphic("skeleton_move_up_8")
	SkeletonMoveUp9 = Graphic("skeleton_move_up_9")
	// Skeleton Move Left Sprites.
	SkeletonMoveLeft1 = Graphic("skeleton_move_left_1") // idle
	SkeletonMoveLeft2 = Graphic("skeleton_move_left_2")
	SkeletonMoveLeft3 = Graphic("skeleton_move_left_3")
	SkeletonMoveLeft4 = Graphic("skeleton_move_left_4")
	SkeletonMoveLeft5 = Graphic("skeleton_move_left_5")
	SkeletonMoveLeft6 = Graphic("skeleton_move_left_6")
	SkeletonMoveLeft7 = Graphic("skeleton_move_left_7")
	SkeletonMoveLeft8 = Graphic("skeleton_move_left_8")
	SkeletonMoveLeft9 = Graphic("skeleton_move_left_9")
	// Skeleton Move Down Sprites.
	SkeletonMoveDown1 = Graphic("skeleton_move_down_1") // idle
	SkeletonMoveDown2 = Graphic("skeleton_move_down_2")
	SkeletonMoveDown3 = Graphic("skeleton_move_down_3")
	SkeletonMoveDown4 = Graphic("skeleton_move_down_4")
	SkeletonMoveDown5 = Graphic("skeleton_move_down_5")
	SkeletonMoveDown6 = Graphic("skeleton_move_down_6")
	SkeletonMoveDown7 = Graphic("skeleton_move_down_7")
	SkeletonMoveDown8 = Graphic("skeleton_move_down_8")
	SkeletonMoveDown9 = Graphic("skeleton_move_down_9")
	// Skeleton Move Right Sprites.
	SkeletonMoveRight1 = Graphic("skeleton_move_right_1") // idle
	SkeletonMoveRight2 = Graphic("skeleton_move_right_2")
	SkeletonMoveRight3 = Graphic("skeleton_move_right_3")
	SkeletonMoveRight4 = Graphic("skeleton_move_right_4")
	SkeletonMoveRight5 = Graphic("skeleton_move_right_5")
	SkeletonMoveRight6 = Graphic("skeleton_move_right_6")
	SkeletonMoveRight7 = Graphic("skeleton_move_right_7")
	SkeletonMoveRight8 = Graphic("skeleton_move_right_8")
	SkeletonMoveRight9 = Graphic("skeleton_move_right_9")
)

// SkeletonMoveUpFrames returns the frames for the SkeletonMoveUp animation.
func SkeletonMoveUpFrames() []Graphic {
	return []Graphic{
		SkeletonMoveUp2,
		SkeletonMoveUp3,
		SkeletonMoveUp4,
		SkeletonMoveUp5,
		SkeletonMoveUp6,
		SkeletonMoveUp7,
		SkeletonMoveUp8,
		SkeletonMoveUp9,
	}
}

// SkeletonMoveLeftFrames returns the frames for the SkeletonMoveLeft animation.
func SkeletonMoveLeftFrames() []Graphic {
	return []Graphic{
		SkeletonMoveLeft2,
		SkeletonMoveLeft3,
		SkeletonMoveLeft4,
		SkeletonMoveLeft5,
		SkeletonMoveLeft6,
		SkeletonMoveLeft7,
		SkeletonMoveLeft8,
		SkeletonMoveLeft9,
	}
}

// SkeletonMoveDownFrames returns the frames for the SkeletonMoveDown animation.
func SkeletonMoveDownFrames() []Graphic {
	return []Graphic{
		SkeletonMoveDown2,
		SkeletonMoveDown3,
		SkeletonMoveDown4,
		SkeletonMoveDown5,
		SkeletonMoveDown6,
		SkeletonMoveDown7,
		SkeletonMoveDown8,
		SkeletonMoveDown9,
	}
}

// SkeletonMoveRightFrames returns the frames for the SkeletonMoveRight animation.
func SkeletonMoveRightFrames() []Graphic {
	return []Graphic{
		SkeletonMoveRight2,
		SkeletonMoveRight3,
		SkeletonMoveRight4,
		SkeletonMoveRight5,
		SkeletonMoveRight6,
		SkeletonMoveRight7,
		SkeletonMoveRight8,
		SkeletonMoveRight9,
	}
}
