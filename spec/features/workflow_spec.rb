%w( capybara capybara/rspec capybara-webkit ).each do |lib|
  require lib
end

include Capybara::DSL

describe 'using the todo app' do
  before do
    Capybara.configure do |config|
      config.run_server = false
      config.current_driver = :webkit
      config.app_host = 'http://localhost:8080'
      config.javascript_driver = :webkit
    end
  end

  describe 'visiting the homepage' do

    it 'allows the user to visit the homepage' do
      visit '/'

      expect(page).to have_content 'Hello world'
    end

    it 'allows the user to see a list of todo items' do
      visit '/'

      expect(page).to have_css 'div#todo_items'

      50.times do |i|
        expect(page).to have_content "Walk Cat #{i}"
      end
      expect(page).to have_content 'Walk Dog'
      expect(page).to have_content 'Walk Giraffe'
      expect(page).to have_content 'Walk Zebra'
      expect(page).to have_content 'Walk Elephant'
    end
  end

  describe 'viewing a todo item' do
    it 'allows the user to view a single todo item' do
      visit '/'

      click_on 'View 0'

      expect(page.current_path).to eq '/view/0'

      expect(page).to have_content "You are viewing 'Walk Cat 0'"
      expect(page).to_not have_content 'Walk Cat 1'
    end
  end
end