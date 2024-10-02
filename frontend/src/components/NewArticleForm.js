import React, { useState } from 'react';

function NewArticleForm() {
    const [name, setName] = useState('');
    const [price, setPrice] = useState(0.0);
    const [wears, setWears] = useState(0);
    const [material, setMaterial] = useState('')
    const [brand, setBrand] = useState('')
    const [season, setSeason] = useState('')
    const [tags, setTags] = useState([])
    const [type, setType] = useState('')

    function addNewArticle(ev) {
      ev.preventDefault();
      const newArticle = {
        name,
        price,
        wears,
        material,
        brand,
        season,
        tags,
        type,
      };
        console.log("New Article:", newArticle);

        setName('');
        setPrice(0.0);
        setWears(0);
        setMaterial('');
        setBrand('');
        setSeason('');
        setTags([]);
        setType('');
    }

    const handleTagsChange = (event) => {
      const options = event.target.options;
      const selectedTags =[];
      for (let i = 0; i < options.length; i++) {
        if (options[i].selected) {
          selectedTags.push(options[i].value)
        }
      }
      setTags(selectedTags)
    }

    return (
      <form onSubmit={addNewArticle}>
        Add More Clothes:
        <div className='first'>
          <label>
            Name:
            <input
              type="text" 
              value={name}
              onChange={ev => setName(ev.target.value)}
              placeholder={'New shirt'}
            />
          </label>
          <label>
            Price:
            <input
              type="number"
              step="0.01"  // To handle float values for price
              value={price}
              onChange={ev => setPrice(ev.target.value)}
              placeholder="0.00"
            />
          </label>
          <label>
            Wears:
            <input
              type="number"
              value={wears}
              onChange={ev => setWears(ev.target.value)}
              placeholder={0}
            />
          </label> 
        </div>
        <div className='second'>
          <label>
            Material:
            <input
              type="text"
              value={material}
              onChange={ev => setMaterial(ev.target.value)}
              placeholder={'Polyester'}
            />
          </label>
          <label>
            Brand:
            <input
              type="text"
              value={brand}
              onChange={ev => setBrand(ev.target.value)}
              placeholder={'Nike'}
            />
          </label>
          <label>
            Season:
            <select
              value={season}
              onChange={ev => setSeason(ev.target.value)}
            >
              <option value="summer">Summer</option>
              <option value="fall">Fall</option>
              <option value="winter">Winter</option>
              <option value="spring">Spring</option>
              
            </select>
          </label>
        </div>
        <div className='labels'>
          <label>
            Tags:
            <select
              multiple
              value={tags}
              onChange={handleTagsChange}
            >
              {/* Options for tags can be added here */}
              <option value="casual">Casual</option>
              <option value="formal">Formal</option>
              <option value="sport">Sport</option>
              
            </select>
            
          </label>
          <label>
            Article Type:
            <select
              value={type}
              onChange={ev => setType(ev.target.value)}
            >
              <option value="top">Top</option>
              <option value="bottom">Bottom</option>
              <option value="shoe">Shoes</option>
              <option value="accessory">Accessories</option>
            
            </select>
            
          </label>
        </div>
        <button type="submit">Submit</button>
      </form>
    );
  }

export default NewArticleForm;