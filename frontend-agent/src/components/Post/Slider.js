import React, { useState } from 'react'

import './../../style/Slider.css';
import { ReactComponent as ChevronRight } from './../../images/icons/chevron-right-solid.svg';
import { ReactComponent as ChevronLeft } from './../../images/icons/chevron-left-solid.svg';
import { ReactComponent as Circle } from './../../images/icons/circle.svg';
import { ReactComponent as Tag } from './../../images/icons/tag.svg';

const Slider = (props) => {
    const { media, showStoryCaption, storyCaption } = props;
    const [slideCount, setSlideCount] = useState(1)
    const [showTags, setShowTags] = useState(props.showTags);

    const nextImage = () => {
        setSlideCount((slideCount + 1 <= media.length) ? slideCount + 1 : slideCount)
    }

    const previousImage = () => {
        setSlideCount((slideCount - 1 >= 1) ? slideCount - 1 : slideCount)
    }

    const Tags = (props) => {
        const { photo } = props;
        return (
            <div className="tagWrapper">
                <Tag onClick={() => setShowTags(!showTags)} className="tagIcon" />
                <div className="tags">
                    {showTags && photo.tags.length > 0 && photo.tags.map(tag => {
                        return (
                            <div className="tag">
                                <div className="tagText">{tag.username}</div>
                            </div>
                        )
                    })}
                </div>
            </div>
        )
    }

    const BackArrow = () => <ChevronLeft onClick={previousImage} className="Arrow-back" />
    const NextArrow = () => <ChevronRight onClick={nextImage} className="Arrow-next" />
    const CircleWrapper = () => {
        if(media.length < 2) return null;

        return (
            <div className="CircleWrapper">
                { media.map((photo) => {
                    const checkActive = photo.orderNum === slideCount ? "active" : "";
                    return (
                        <Circle className={`Circle ${checkActive}`} />
                    )
                })}
            </div>
        )}
        
    const StoryCaption = () => (
        <div className="storyCaptionWrapepr">
            <div className="storyCaptionText">{storyCaption}</div>
        </div>
    )

    return (
        <div className="Slide">
            {media.map((photo) => {
                if (photo.orderNum === slideCount) {
                    return (
                        <div key={photo.id}>
                            <img src={photo.content} alt=''/>
                            <div className="tags-caption">
                                { photo.tags.length !== 0 && <Tags photo={photo} /> }
                                { showStoryCaption && <StoryCaption /> }
                            </div>
                        </div>
                    )
                }
                return ''
            })}

            <CircleWrapper />
            {slideCount !== 1 ? <BackArrow /> : ''}
            {slideCount !== media.length ? <NextArrow /> : ''}
        </div>
    );
}

export default Slider;